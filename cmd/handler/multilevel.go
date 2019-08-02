package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"multi-level-marketing-project/internal/member"
	"net/http"
	"strconv"
)

func GetMember(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	memberID, err := strconv.Atoi(vars["memberID"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	member := member.FindMember(memberID)
	encode := json.NewEncoder(writer)
	err = encode.Encode(&member)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
