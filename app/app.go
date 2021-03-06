package app

import (
	"net/http"
	"text/template"

	"github.com/GibranHL0/devblog/configuration"
	"github.com/GibranHL0/devblog/connection"
)

type App struct {
	Config    *configuration.Configuration
	Db        *connection.Database
	Templates *template.Template
	Server    *http.ServeMux
}

func Factory(
	config *configuration.Configuration,
	db *connection.Database,
	templates *template.Template,
	server *http.ServeMux) App {

	return App{
		Config: config,
		Db: db,
		Templates: templates,
		Server: server,
	}
}
