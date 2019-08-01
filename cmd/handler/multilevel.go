package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"multi-level-marketing-project/database/db"
	"multi-level-marketing-project/internal/members"
	"multi-level-marketing-project/internal/point"
	"multi-level-marketing-project/models"
	"net/http"
)

func AddPoint(w http.ResponseWriter, r *http.Request) {
	newUserPoint, err := decodeRequest(r)
	if err != nil {
		http.Error(w, "Error Decode request unsuccess", http.StatusBadRequest)
		return
	}
	db, err := db.DBConnection()
	if err != nil {
		http.Error(w, "Error Database connection unsuccess", http.StatusInternalServerError)
		return
	}
	if point.RecordPoint(db, newUserPoint) == false {
		http.Error(w, "Error Record point unsuccess", http.StatusBadRequest)
		return
	} else {
		members.VerifyLevel(db, newUserPoint.UserRefferal)
	}
}

func decodeRequest(r *http.Request) (models.NewUserPoint, error) {
	var newUserPoint models.NewUserPoint
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
