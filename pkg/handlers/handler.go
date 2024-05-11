package handlers

import (
	"net/http"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate("../templates/layout.html", w)

}
