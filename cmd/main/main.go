package main

import (
	"log"
	"multi-level-marketing-project/cmd/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/new_user_point", handler.AddPoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
