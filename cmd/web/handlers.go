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
		"ui/html/partials/header.html",
	}

	tmpl,err := template.ParseFiles(files...)

	if err!=nil{
		app.internalError(w,r,err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", statuses)
	if err!=nil{
		app.internalError(w,r,err)
		return
	}
}

func (app *application) status(w http.ResponseWriter, r *http.Request){
	// id := r.PathValue("codeId")

	files := []string{
		"ui/html/pages/status.html",
		"ui/html/templates/base.tmpl.html",
		"ui/html/partials/header.html",
	}

	tmpl,err := template.ParseFiles(files...)

	if err!=nil{
		app.internalError(w,r,err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", statuses)
	if err!=nil{
		app.internalError(w,r,err)
		return
	}

}

var statuses = []Status{
	{100, "Continue"},
	{101, "Switching Protocols"},
	{102, "Processing"},
	{200, "OK"},
	{201, "Created"},
	{204, "No Content"},
	{301, "Moved Permanently"},
	{302, "Found"},
	{304, "Not Modified"},
	{400, "Bad Request"},
	{401, "Unauthorized"},
	{403, "Forbidden"},
	{404, "Not Found"},
	{418, "I'm a teapot"},
	{429, "Too Many Requests"},
	{500, "Internal Server Error"},
	{501, "Not Implemented"},
	{502, "Bad Gateway"},
	{503, "Service Unavailable"},
	{504, "Gateway Timeout"},
}