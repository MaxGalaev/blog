package main

import (
	"net/http"

	"github.com/MaxGalaev/blog/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.MyHandler)
	http.ListenAndServe(":8080", nil)
}
