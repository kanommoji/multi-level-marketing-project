package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql",
		"root:root@tcp(10.10.99.37:3306)/mlm")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
