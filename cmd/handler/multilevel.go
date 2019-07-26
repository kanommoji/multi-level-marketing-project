package handler

import (
	"multi-level-marketing-project/internal/members"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"multi-level-marketing-project/internal/point"
)

func AddPoint(w http.ResponseWriter, r *http.Request) {
	pointAction, err := decodeAction(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !point.RecordPoint(pointAction) {
		w.WriteHeader(http.StatusBadRequest)
	}
	members.VerifyLevel(pointAction.UserReferral)
}

func decodeAction(request *http.Request) (point.Action, error) {
	var action point.Action
	body, err := ioutil.ReadAll(request.Body)
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
	vars := mux.Vars(r)
	memberId, err := strconv.Atoi(vars["key"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	member := members.FindMember(memberId)
	err = json.NewEncoder(w).Encode(&member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
