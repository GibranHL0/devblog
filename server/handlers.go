package server

import (
	"net/http"
	"text/template"
)

func aboutHandler(templates *template.Template) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		executeTemplate(templates, "about.html", nil, rw)
	}
}

func contactHandler(templates *template.Template) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		executeTemplate(templates, "contact.html", nil, rw)
	}
}

func newsletterHandler(templates *template.Template) http.HandlerFunc{
	return func(rw http.ResponseWriter, r *http.Request) {
		executeTemplate(templates, "newsletter.html", nil, rw)
	}
}