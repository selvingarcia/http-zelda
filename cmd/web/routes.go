package main

import (
	"net/http"
	"strings"
	"github.com/selvingarcia/http-zelda/ui"
)
func (app *application) routes() *http.ServeMux{
	mux := http.NewServeMux()
	mux.Handle("GET /static/", neuter(http.FileServerFS(ui.Files)))
	mux.HandleFunc("GET /", app.home)
	mux.HandleFunc("GET /status/{codeId}", app.status)
	mux.HandleFunc("GET /search", app.search)
	mux.HandleFunc("GET /not-found", app.notFoundSearch)
	return mux
}

// eliminate static path 
func neuter(file http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/"){
			http.NotFound(w,r)
			return
		}

		file.ServeHTTP(w,r)
	})
}