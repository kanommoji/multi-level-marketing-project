package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"multi-level-marketing-project/config"
	"multi-level-marketing-project/database"
	"multi-level-marketing-project/internal/member"
	"multi-level-marketing-project/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetMember(writer http.ResponseWriter, request *http.Request) {
	config := config.Config{
		Username: "root",
		Password: "root",
		Host:     "127.0.0.1",
		Database: "mlm",
		Port:     "3306",
	}
	database, err := database.DBConnect(config.GetURI())

	vars := mux.Vars(request)
	memberID, err := strconv.Atoi(vars["memberID"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	member := member.FindMember(database, memberID)
	encode := json.NewEncoder(writer)
	err = encode.Encode(&member)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func decodeNewUserPoint(r *http.Request) (model.NewUserPoint, error) {
	var newUserPoint model.NewUserPoint
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return newUserPoint, err
	}
	decode := json.NewDecoder(bytes.NewReader(body))
	err = decode.Decode(&newUserPoint)
	if err != nil {
		return newUserPoint, err
	}
	return newUserPoint, nil
}
