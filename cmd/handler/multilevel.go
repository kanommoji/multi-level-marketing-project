package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"multi-level-marketing-project/internal/member"
	"multi-level-marketing-project/internal/point"
	"multi-level-marketing-project/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DatabaseConnection struct {
	SQLDatabase *sql.DB
}

func (database DatabaseConnection) GetMember(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	memberID, err := strconv.Atoi(vars["memberID"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	member := member.FindMember(database.SQLDatabase, memberID)
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

func (database DatabaseConnection) AddPoint(writer http.ResponseWriter, request *http.Request) {
	newUserPoint, err := decodeNewUserPoint(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	if point.RecordPoint(database.SQLDatabase, newUserPoint) == false {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	} else {
		member.VerifyLevel(database.SQLDatabase, newUserPoint.UserReferral)
	}
}
