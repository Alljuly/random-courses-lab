package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//String de conexão com o banco
	StringSQLConn = ""

	//Porta da API
	Port = 0
)

func Prepare() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 9000
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_SENHA")
	db := os.Getenv("DB_NOME")

	StringSQLConn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=true&loc=Local", user, pass, db)
}