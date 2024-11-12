package models

import "cadastroProduto/db"

type Produto struct{
	Id int
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

func GetAllProdutos()[]Produto{
	banco := db.ConectaBanco()

	produtos := []Produto{}
	p := Produto{}

	all, err := banco.Query("SELECT * FROM tab_produtos")

	if err != nil{
		panic(err.Error())
	}

	for all.Next(){
		var id, quantidade int
		var nome, descricao string
		var preco float64

		all.Scan(&id, &nome, &descricao, &preco, &quantidade)

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer banco.Close()
	return produtos
}