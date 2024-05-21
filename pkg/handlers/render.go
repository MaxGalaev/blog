package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type BlogPost struct {
	Title   string
	Content string
}

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

func ParseMarkdown(MarkdownPath string) BlogPost {
	markdownfile, err := os.ReadFile(MarkdownPath)
	if err != nil {
		log.Println(err.Error())
	}
	title := string(markdownfile[0])
	content := strings.Split(string(markdownfile), "___")
	blogpost := BlogPost{Title: title, Content: content[1]}
	return blogpost

}
