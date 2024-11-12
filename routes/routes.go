package routes

import(
	"net/http"
	"cadastroProduto/controllers"
)

func CarregaRotas(){
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.InsertScreen)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
}