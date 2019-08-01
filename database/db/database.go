package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql",
		"root:root@tcp(10.10.99.37:3306)/mlm")
	if err != nil {
		return nil, err
	}
	return db, nil
}
