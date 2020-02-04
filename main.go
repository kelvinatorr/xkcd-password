package main

import (
	"bufio"
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type fileSystem struct {
	fs http.FileSystem
}

// Open opens file if it exists in the directory. If path is a directory it checks if index.html exists
// and returns nil, err if it does not.
func (fs fileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		// Check if index.html exists. If it does not return 404.
		// Without this, the default behavior is to return a directory listing, which we don't want.
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}
	return f, nil
}

func main() {
	rawWordsPath := flag.String("raw-words-path", "/etc/xkcd-password/words-raw.txt",
		"The path to a list of common english words. One on each line")
	flag.Parse()

	// Check that the word file is ready
	preparedFilePath := "words.txt"
	if _, err := os.Stat(preparedFilePath); os.IsNotExist(err) {
		prepareWordFile(*rawWordsPath, preparedFilePath)
	}

	log.Println("Starting webserver. Listening on port 8080")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", func(w http.ResponseWriter, req *http.Request) {
		var data interface{}
		t, _ := template.ParseFiles("templates/about.html")
		t.Execute(w, data)
	})
	// Handle files in the static directory
	fs := http.FileServer(fileSystem{http.Dir("static")})
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe("localhost:8080", nil)
}

func checkAndLog(err error) {
	if err != nil {
		log.Println(err)
	}
}

func checkAndPanic(err error) {
	if err != nil {
		// Panic but don't exit
		panic(err)
	}
}

func prepareWordFile(rawFilePath string, preparedFilePath string) {

	scanner, file := getWordFileScanner(rawFilePath)
	defer file.Close()

	f, err := os.Create(preparedFilePath)
	checkAndPanic(err)

	w := bufio.NewWriter(f)

	for scanner.Scan() {
		s := scanner.Text()
		_, err := w.WriteString(strings.TrimSpace(s) + "\n")
		checkAndLog(err)
	}

	w.Flush()

	err = scanner.Err()
	checkAndLog(err)
}

// getWordFileScanner returns a bufio.Scanner for the given wordFilePath. Also returns the open os.File
// object which the calling function should Close()
func getWordFileScanner(wordFilePath string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(wordFilePath)
	checkAndPanic(err)

	return bufio.NewScanner(file), file
}
