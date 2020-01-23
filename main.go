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
		log.Fatal(err)
	}
}

func checkAndPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func prepareWordFile(rawFilePath string, preparedFilePath string) {
	file, err := os.Open(rawFilePath)
	checkAndPanic(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

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
