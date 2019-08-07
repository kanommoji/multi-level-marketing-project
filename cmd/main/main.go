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
	databaseConfig, err := config.SetupConfig()
	if err != nil {
		log.Fatalf("Setup config error %s", err.Error())
	}
	databaseConnection, err := database.DBConnect(databaseConfig.GetURI())
	if err != nil {
		log.Fatalf("Database connect failed %s", err.Error())
	}
	multilevelHandler := handler.DatabaseConnection{
		SQLDatabase: databaseConnection,
	}

	router := mux.NewRouter()
	router.HandleFunc("/new_user_point", multilevelHandler.AddPoint).Methods("POST")
	router.HandleFunc("/members/{memberID}", multilevelHandler.GetMember).Methods("GET")
	router.HandleFunc("/verify_demote", multilevelHandler.DemoteMember).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
