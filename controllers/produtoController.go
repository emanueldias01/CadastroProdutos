package controllers

import (
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*html"))

func Index(w http.ResponseWriter, r *http.Request){
	temp.ExecuteTemplate(w, "Index", nil)
}