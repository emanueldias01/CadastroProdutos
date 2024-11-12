package main

import (
	"cadastroProduto/routes"
	"net/http"
)

func main(){
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}