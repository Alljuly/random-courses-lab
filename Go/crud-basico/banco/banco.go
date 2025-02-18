package banco

import (

	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o MySQL
)

//Conectar ao banco
func Connection() (*sql.DB, error){
	conn :=  "golang:golang@/devbook?charset=utf8&parseTime=true&loc=Local"

	db, erro := sql.Open("mysql", conn)


	if erro != nil{
		return nil, erro
	}

	if erro = db.Ping(); erro != nil{
		return nil, erro
	}

	return db, nil
}