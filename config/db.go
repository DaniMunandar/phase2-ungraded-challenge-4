package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/phase2_ngc4")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}
