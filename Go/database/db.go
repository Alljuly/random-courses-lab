package main

import (
	"fmt"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	conn :=  "golang:golang@/devbook?charset=utf8&parseTime=Time&loc=Local"

	db, erro := sql.Open("mysql", conn)

	if erro != nil {
		log.Fatal(erro)
	}


	fmt.Println(db)

}