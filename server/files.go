package server

import "net/http"

func createFileServer(path string) http.Handler {
	return http.FileServer(http.Dir(path))
}