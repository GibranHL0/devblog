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
