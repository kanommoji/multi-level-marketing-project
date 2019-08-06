package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnect(url string) (*sql.DB, error) {
	database, err := sql.Open("mysql", url)
	if err != nil {
		return database, err
	}
	return database, nil
}
