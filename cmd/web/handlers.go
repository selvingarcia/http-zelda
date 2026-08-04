package main

import(
	"net/http"
	"html/template"
)

//main webpage 

func (app *application) home(w http.ResponseWriter, r *http.Request){
	files := []string{
		"ui/html/pages/home.html",
		"ui/html/templates/base.tmpl.html",
	}

	tmpl,err := template.ParseFiles(files...)

	if err!=nil{
		app.internalError(w,r,err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err!=nil{
		app.internalError(w,r,err)
		return
	}
}