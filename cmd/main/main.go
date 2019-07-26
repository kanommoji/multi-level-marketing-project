package main

import (
	"log"
	"multi-level-marketing-project/cmd/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/action_point", handler.AddPoint).Methods("POST")
	router.HandleFunc("/members/{key}", handler.GetMember).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
