package main

import (
	"net/http"
)
func (app *application) routes() *http.ServeMux{
	fs := http.FileServer(http.Dir("./ui/static"))
	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.StripPrefix("/static",neuter(fs)))
	mux.HandleFunc("GET /", app.home)
	mux.HandleFunc("GET /status/{codeId}", app.status)
	return mux
}
	