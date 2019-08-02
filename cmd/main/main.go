package main

import (
	"github.com/gorilla/mux"
	"log"
	"multi-level-marketing-project/cmd/handler"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/members/{memberID}", handler.GetMember).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
