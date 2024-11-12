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

func CadastraProduto(nome string, descricao string, preco float64, quantidade int){
	banco := db.ConectaBanco()

	insert, err := banco.Prepare("INSERT INTO tab_produtos(nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)")

	if err != nil{
		panic(err.Error())
	}

	insert.Exec(nome, descricao, preco, quantidade)

	defer banco.Close()
}