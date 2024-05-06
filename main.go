package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
