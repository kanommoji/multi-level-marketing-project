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

func (databaseConnection DatabaseConnection) GetMember(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	memberID, err := strconv.Atoi(vars["memberID"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	member := member.FindMember(databaseConnection.SQLDatabase, memberID)
	encode := json.NewEncoder(writer)
	err = encode.Encode(&member)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func decodeNewUserPoint(request *http.Request) (model.NewUserPoint, error) {
	var newUserPoint model.NewUserPoint
	body, err := ioutil.ReadAll(request.Body)
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

func (databaseConnection DatabaseConnection) AddPoint(writer http.ResponseWriter, request *http.Request) {
	newUserPoint, err := decodeNewUserPoint(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	if !point.RecordPoint(databaseConnection.SQLDatabase, newUserPoint) {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	member.VerifyLevel(databaseConnection.SQLDatabase, newUserPoint.UserReferral)
}

func decodeMemberID(request *http.Request) (model.Member, error) {
	var member model.Member
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return member, err
	}
	decode := json.NewDecoder(bytes.NewReader(body))
	err = decode.Decode(&member)
	if err != nil {
		return member, err
	}
	return member, nil
}

func (databaseConnection DatabaseConnection) DemoteMember(writer http.ResponseWriter, request *http.Request) {
	memberDemote, err := decodeMemberID(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	member.VerifyLevelDemote(databaseConnection.SQLDatabase, memberDemote.MemberID)
}
