package main

import (
	"log"
	"net/http"

	"multi-level-marketing-project/cmd/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/action_point", handler.AddPoint).Methods("POST")
	r.HandleFunc("/members/{key}", handler.GetMember).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
