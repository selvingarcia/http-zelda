package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/selvingarcia/http-zelda/internal/models"
	"github.com/selvingarcia/http-zelda/ui"
)

//main webpage



func (app *application) home(w http.ResponseWriter, r *http.Request){
	files := []string{
		"html/pages/home.html",
	}

	baseFiles := baseTemplate(BaseTemplate, files)

	tmpl,err := template.ParseFS(ui.Files,baseFiles...)

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
	id,err := toIntN(r.PathValue("codeId"))
	if err!=nil{
		app.badRequest(w,r,err)
		return
	}


	files := []string{
		"html/pages/status.html",
	}

	baseFiles := baseTemplate(BaseTemplate, files)

	tmpl,err := template.ParseFS(ui.Files,baseFiles...)

	if err!=nil{
		app.internalError(w,r,err)
		return
	}

	status, err := app.db.GetStatus(id)

	if err!=nil{
		if errors.Is(err, models.NoRows){
			app.notFound(w,r,err)
			return
		}else{
			app.internalError(w,r,err)
			return
		}
	}

	data := StatusData{
		Status: status,
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err!=nil{
		app.internalError(w,r,err)
		return
	}

}

func (app *application) search(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query().Get("q")
	v,err := toIntN(query)
	if err!=nil{
		http.Redirect(w,r,"/not-found",http.StatusSeeOther)
		return
	}
	for _,n := range(statuses){
		if n.Code == v{
			f := fmt.Sprintf("/status/%d", v)
			http.Redirect(w,r,f,http.StatusSeeOther)
			return
		}
	}

	http.Redirect(w, r, "/not-found", http.StatusSeeOther)
}

func (app *application) notFoundSearch(w http.ResponseWriter, r *http.Request){
		files := []string{
		"html/pages/not-found.html",
	}

	baseFiles := baseTemplate(BaseTemplate, files)

	tmpl,err := template.ParseFS(ui.Files,baseFiles...)

	if err!=nil{
		app.internalError(w,r,err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", baseFiles)
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

func toIntN(n string) (int, error){
	number , err :=strconv.Atoi(n)
	if err!=nil{
		return 0, err
	}
	return number, nil
}

func baseTemplate(baseArray , addArray []string) []string{
	b:= make([]string,0, len(baseArray)+len(addArray))
	b = append(b, baseArray...)
	b = append(b, addArray...)
	return b
}