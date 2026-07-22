package main

import(
	"net/http"
	"html/template"
)


func (app *application) home(w http.ResponseWriter, r *http.Request){
	files := []string{
		"ui/html/home.html",
	}

	tmpl,err := template.ParseFiles(files...)

	if err!=nil{
		app.logger.Error(err.Error(), "Method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

	tmpl.ExecuteTemplate(w, "home.html", nil)
}