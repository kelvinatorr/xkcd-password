package main

import (
	"bufio"
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
)

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
	http.HandleFunc("/", IndexHandler)
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

	scanner, file := GetWordFileScanner(rawFilePath)
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

// GetWordFileScanner returns a bufio.Scanner for the given wordFilePath. Also returns the open os.File
// object which the calling function should Close()
func GetWordFileScanner(wordFilePath string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(wordFilePath)
	checkAndPanic(err)

	return bufio.NewScanner(file), file
}
