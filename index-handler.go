package main

import (
	"fmt"
	"net/http"
)

// IndexHandler Handler for the / path
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yay index")
}
