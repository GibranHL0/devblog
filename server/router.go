package server

import (
	"net/http"
	"text/template"

	"github.com/GibranHL0/devblog/configuration"
)

func initializeHandlers(
	mux *http.ServeMux,
	config *configuration.Configuration,
	templates *template.Template) {
	
	// Creates a http handler that will work as a File Server
	fileServer := createFileServer(config.FileServerPath)

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/about", aboutHandler(templates))
}

func StartMux(
	config *configuration.Configuration,
	templates *template.Template) *http.ServeMux {
	// Starts a fresh Mux
	mux := http.NewServeMux()

	// Routes all the handlers with their respective pattern
	initializeHandlers(mux, config, templates)

	return mux
}
