package server

import (
	"net/http"
	"text/template"

	"github.com/GibranHL0/devblog/errorhandler"
)

func aboutHandler(templates *template.Template) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(rw, "about.html", "")
		if err != nil {
			errorhandler.ReportError(err, "about.html was not executed.")
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Something bad happened!"))
		}
	}
}
