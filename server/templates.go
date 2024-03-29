package server

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseFiles("./server/templates/index.tmpl", "./server/templates/dashboard.tmpl"))

// RenderTemplate : render given template to response writer
func RenderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".tmpl", p)
	if err != nil {
		log.Printf("Temlate error here: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
