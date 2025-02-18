package database

import ("api/src/config"
		"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringSQLConn)

	if err != nil{
		return nil, err
	}

	if err = db.Ping(); err != nil{
		db.Close()
	}

	return db, nil
}
