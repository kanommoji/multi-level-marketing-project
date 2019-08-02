package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnect(url string) (*sql.DB, error) {
	database, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Println("database connect fail")
		return database, err
	}
	return database, nil
}
