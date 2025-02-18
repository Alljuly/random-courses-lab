package main

import (
	"fmt"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	conn :=  "golang:golang@/devbook?charset=utf8&parseTime=true&loc=Local"

	db, erro := sql.Open("mysql", conn)

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conn aberta")
	
	
	linhas, erro := db.Query("Select * from users")
	
	if erro != nil{
		log.Fatal(erro)
	}
	fmt.Println(linhas)
	defer linhas.Close()

}