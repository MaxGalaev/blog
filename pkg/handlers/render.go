package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(templatepath string, w http.ResponseWriter) error {
	// parse html template
	post_template, err := template.ParseFiles(templatepath)
	if err != nil {
		log.Println(err.Error())
	}
	// execute template
	err = post_template.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
	}
	return err
}
