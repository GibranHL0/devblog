package server

import (
	"net/http"

	"github.com/GibranHL0/devblog/configuration"
)

func initializeHandlers(mux *http.ServeMux, fs http.Handler) *http.ServeMux {
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	return mux
}

func StartMux(config configuration.Configuration) *http.ServeMux {
	// Starts a fresh Mux
	mux := http.NewServeMux()

	// Creates a http handler that will work as a File Server
	fileServer := createFileServer(config.FileServerPath)

	// Routes all the handlers with their respective pattern
	initializeHandlers(mux, fileServer)
	
	return mux
}