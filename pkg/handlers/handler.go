package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	// parse html template
	post_template, err := template.ParseFiles("../templates/layout.html")
	if err != nil {
		log.Println(err.Error())
	}
	// execute template
	post_template.Execute(w, nil)
}
