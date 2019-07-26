package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"multi-level-marketing-project/internal/members"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"multi-level-marketing-project/internal/point"
)

func AddPoint(w http.ResponseWriter, r *http.Request) {
	action, err := decodeAction(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !point.RecordPoint(action) {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		members.VerifyLevel(action.UserReferral)
	}
}

func decodeAction(r *http.Request) (point.Action, error) {
	var action point.Action
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return action, err
	}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&action)
	if err != nil {
		return action, err
	}
	return action, nil
}

func GetMember(w http.ResponseWriter, r *http.Request) {
	pathVariables := mux.Vars(r)
	memberID, err := strconv.Atoi(pathVariables["key"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	member := members.FindMember(memberID)
	err = json.NewEncoder(w).Encode(&member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
