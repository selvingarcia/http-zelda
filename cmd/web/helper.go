package main

import "net/http"

func (app *application) internalError(w http.ResponseWriter,r *http.Request, err error){
	app.logger.Error(err.Error(), "Method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
}