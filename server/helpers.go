package server

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/GibranHL0/devblog/errorhandler"
	"github.com/GibranHL0/devblog/models"
)

func executeTemplate(
	templates *template.Template, 
	name string, 
	data models.IModel,
	rw http.ResponseWriter){
	
	err := templates.ExecuteTemplate(rw, name, data)
	if err != nil {
		info := fmt.Sprintf("%s template was not executed", name)
		errorhandler.ReportError(err, info)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Something bad happened!"))
	}
}