package main

import (
	"log"
	"multi-level-marketing-project/cmd/handler"
	"multi-level-marketing-project/config"
	"multi-level-marketing-project/database"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config, _ := config.SetupConfig()
	db, _ := database.DBConnect(config.GetURI())
	multilevelHandler := handler.DatabaseConnection{
		SQLDatabase: db,
	}

	router := mux.NewRouter()
	router.HandleFunc("/new_user_point", multilevelHandler.AddPoint).Methods("POST")
	router.HandleFunc("/members/{memberID}", multilevelHandler.GetMember).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
