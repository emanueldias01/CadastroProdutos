package main

import(
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*html"))

func main(){
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, r *http.Request){
	temp.ExecuteTemplate(w, "Index", nil)
}