package main

import (
	"net/http"
	"os"
	"text/template"
)

type aboutPage struct {
	WordCount int
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(preparedFilePath)
	checkAndPanic(err)
	var data aboutPage = aboutPage{}
	data.WordCount, err = lineCounter(f)
	checkAndLog(err)
	t, _ := template.ParseFiles("templates/about.html")
	t.Execute(w, data)
}
