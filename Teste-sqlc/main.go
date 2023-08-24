package main

import (
	"context"
	"database/sql"
	db "main/gen"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dbcon, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/BaseParaTestes")
	if err != nil {
		panic(err.Error())
	}
	defer dbcon.Close()

	// Execute the query
	dt := db.New(dbcon)

	ctx := context.Background()

	Pessoa1, err1 := dt.SelectPessoaBycodigo(ctx, 1)

	if err1 != nil {
		panic(err1.Error())
	}

	println("Código	:", Pessoa1.Nome)
	println("E-mail	:", Pessoa1.Email)
	println("idade	:", Pessoa1.Idade)
}
