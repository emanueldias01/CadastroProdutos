package main

import(
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("/templates/*html"))

func main(){

	http.ListenAndServe(":8000", nil)
}