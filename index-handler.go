package main

import (
	"net/http"
	"text/template"
)

// IndexPage struct to fill in the index.html template
type IndexPage struct {
	GeneratedPassword string
}

// IndexHandler Handler for the / path
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		FourZeroFourHandler(w, r, http.StatusNotFound)
		return
	}
	// TODO: Get number of words from the GET params.
	password := GeneratePassword(4)
	indexPage := IndexPage{GeneratedPassword: password}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, indexPage)
}
