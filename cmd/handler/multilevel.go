package handler

import (
	"encoding/json"
	"multi-level-marketing-project/internal/member"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetMember(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	memberID, err := strconv.Atoi(vars["key"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	member := member.FindMember(memberID)
	encode := json.NewEncoder(writer)
	err = encode.Encode(&member)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
