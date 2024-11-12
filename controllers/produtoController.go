package controllers

import (
	"cadastroProduto/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*html"))

func Index(w http.ResponseWriter, r *http.Request){
	temp.ExecuteTemplate(w, "Index", models.GetAllProdutos())
}

func Insert(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoFloat, err := strconv.ParseFloat(preco, 64)

		if err !=nil{
			log.Println("erro na conversao de preco")
		}

		quantidadeInt, err := strconv.Atoi(quantidade)

		if err != nil{
			log.Println("erro na conversao de quantidade")
		}

		models.CadastraProduto(nome, descricao, precoFloat, quantidadeInt)
	}
}