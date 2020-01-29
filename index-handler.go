package main

import (
	"net/http"
	"strconv"
	"text/template"
)

// IndexPage struct to fill in the index.html template
type IndexPage struct {
	GeneratedPassword string
	NumberOfWords     int
}

// IndexHandler Handler for the / path
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		FourZeroFourHandler(w, r, http.StatusNotFound)
		return
	}
	getParams := r.URL.Query()
	numberOfWords := 4
	if requestedNumberOfWordsSlice, ok := getParams["numberOfWords"]; ok {
		requestedNumberOfWords, err := strconv.Atoi(requestedNumberOfWordsSlice[0])
		checkAndLog(err)
		if err == nil && requestedNumberOfWords >= 3 && requestedNumberOfWords <= 8 {
			numberOfWords = requestedNumberOfWords
		}
	}
	// Get number of words from the GET params.
	password := GeneratePassword(numberOfWords)
	indexPage := IndexPage{GeneratedPassword: password, NumberOfWords: numberOfWords}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, indexPage)
}
