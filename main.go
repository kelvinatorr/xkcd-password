package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting webserver. Listening on port 8080")
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe("localhost:8080", nil)
}
