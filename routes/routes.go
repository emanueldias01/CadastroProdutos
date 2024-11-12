package routes

import(
	"net/http"
	"cadastroProduto/controllers"
)

func CarregaRotas(){
	http.HandleFunc("/", controllers.Index)
}