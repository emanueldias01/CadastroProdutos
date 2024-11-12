package db

import(
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaBanco() *sql.DB{

	strConn := "user=postgres password=root dbname=manel_loja sslmode=disable"

	db, err := sql.Open("postgres", strConn)

	if err != nil{
		panic(err.Error())
	}

	return db
}