package main

import "net/http"

func (app *application) internalError(w http.ResponseWriter,r *http.Request, err error){
	app.logger.Error(err.Error(), "Method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error){
	app.logger.Error(err.Error(), "Method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w,"Bad Request", http.StatusBadRequest)
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request, err error){
	app.logger.Error(err.Error(), "Method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w,"Not Found", http.StatusNotFound)
}