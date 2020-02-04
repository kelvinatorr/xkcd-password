package main

import (
	"fmt"
	"net/http"
)

// fourZeroFourHandler handles 404 not found situations.
func fourZeroFourHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404 page not found")
	}
}
