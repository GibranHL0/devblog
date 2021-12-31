package server

import (
	"net/http"
	"text/template"

	"github.com/GibranHL0/devblog/configuration"
	"github.com/GibranHL0/devblog/connection"
)

func initializeHandlers(
	mux *http.ServeMux,
	db *connection.Database,
	templates *template.Template,
	config *configuration.Configuration) {
	
	// Creates a http handler that will work as a File Server
	fileServer := createFileServer(config.FileServerPath)

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", homeHandler(templates, db))
	mux.HandleFunc("/article", articleHandler(templates, db))
	mux.HandleFunc("/about", aboutHandler(templates))
	mux.HandleFunc("/contact", contactHandler(templates))
	mux.HandleFunc("/suscribe", newsletterHandler(templates))
}

func StartMux(
	db *connection.Database,
	config *configuration.Configuration,
	templates *template.Template) *http.ServeMux {
	// Starts a fresh Mux
	mux := http.NewServeMux()

	// Routes all the handlers with their respective pattern
	initializeHandlers(mux, db, templates, config)

	return mux
}
