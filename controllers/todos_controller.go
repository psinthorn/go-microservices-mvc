package controllers

import (
	"net/http"
	"text/template"
)

// var tmpl *template.Template

func TodoIndex(res http.ResponseWriter, req *http.Request) {
	/* Step and Requirment Preparing
	1. Template
	2. Data
	3. Render
	*/

	//Template
	tmpl := template.Must(template.ParseFiles("./views/todos/index.html"))

	tmpl.Execute(res, "Todo Index")
}
