package main

import (
	"html/template"
	"log"
	"net/http"
)

type post struct {
	Title string
	Body  string
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", nil)
}
func myHandler(w http.ResponseWriter, r *http.Request) {
	// parse html template
	post_template, err := template.ParseFiles("./templates/layout.html")
	if err != nil {
		log.Println(err.Error())
	}
	post1 := post{
		Title: "Ahaha",
		Body:  "New post",
	}
	// execute template
	post_template.Execute(w, post1)
}
