package main

import (
	"fmt"
	"net/http"
)

// FourZeroFourHandler handles 404 not found situations.
func FourZeroFourHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404")
	}
}
