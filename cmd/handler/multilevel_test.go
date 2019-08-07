package handler

import (
	"bytes"
	"encoding/json"
	"multi-level-marketing-project/config"
	"multi-level-marketing-project/database"
	"multi-level-marketing-project/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func Test_GetMember_By_MemberID_10029_Should_Get_Member(t *testing.T) {
	var expectedResult, actualResult model.Member
	response := `{"member_id":10029,"member_name":"ชนา","leader_id":20029,"level":6,"my_point":1000,"monthly_point":350,"team_point":20000,"team_member":{"higher_pearl":2,"higher_emerald":2}}`
	decodeExpectedResult := json.NewDecoder(bytes.NewReader([]byte(response)))
	_ = decodeExpectedResult.Decode(&expectedResult)
	request := httptest.NewRequest("GET", "/members/10029", nil)
	request = mux.SetURLVars(request, map[string]string{
		"memberID": "10029",
	})
	request.Header.Set("Content-type", "application/json")
	writer := httptest.NewRecorder()
	configURI := config.Config{
		Username: "mlm_dev",
		Password: "mlm_dev",
		Host:     "127.0.0.1",
		Database: "mlm",
		Port:     "3306",
	}
	databaseConnection, _ := database.DBConnect(configURI.GetURI())
	multilevelHandler := DatabaseConnection{
		SQLDatabase: databaseConnection,
	}

	multilevelHandler.GetMember(writer, request)
	decodeActualResult := json.NewDecoder(writer.Body)
	_ = decodeActualResult.Decode(&actualResult)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_Decode_Action_From_RequestBody_Should_Get_Struct_Action(t *testing.T) {
	expectedResult := model.NewUserPoint{
		NewPoint:     50,
		UserReferral: 10029,
	}
	requestBody := `{"user_referral":10029,"new_point":50}`
	request := httptest.NewRequest("POST", "/new_user_point", bytes.NewReader([]byte(requestBody)))

	actualResult, _ := decodeNewUserPoint(request)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_AddPoint_By_UserReferral_10029_NewPoint_50_Should_Get_StatusOK(t *testing.T) {
	expectedResult := http.StatusOK
	request := httptest.NewRequest("POST", "/new_user_point", strings.NewReader(`{"user_referral":10029,"new_point":50}`))
	writer := httptest.NewRecorder()
	configURI := config.Config{
		Username: "mlm_dev",
		Password: "mlm_dev",
		Host:     "127.0.0.1",
		Database: "mlm",
		Port:     "3306",
	}
	databaseConnection, _ := database.DBConnect(configURI.GetURI())
	multilevelHandler := DatabaseConnection{
		SQLDatabase: databaseConnection,
	}
	multilevelHandler.AddPoint(writer, request)

	actualResult := writer.Code

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_Decode_MemberID_From_RequestBody_Should_Get_Struct_Member(t *testing.T) {
	expectedResult := model.Member{
		MemberID: 40001,
	}
	requestBody := `{"member_id":40001}`
	request := httptest.NewRequest("POST", "/verify_demote", bytes.NewReader([]byte(requestBody)))

	actualResult, _ := decodeMemberID(request)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_DecodeMember_By_MemberID_40001_Should_Get_StatusOK(t *testing.T) {
	expectedResult := http.StatusOK
	request := httptest.NewRequest("POST", "/verify_demote", strings.NewReader(`{"member_id":40001}`))
	writer := httptest.NewRecorder()
	configURI := config.Config{
		Username: "mlm_dev",
		Password: "mlm_dev",
		Host:     "127.0.0.1",
		Database: "mlm",
		Port:     "3306",
	}
	databaseConnection, _ := database.DBConnect(configURI.GetURI())
	multilevelHandler := DatabaseConnection{
		SQLDatabase: databaseConnection,
	}
	multilevelHandler.DemoteMember(writer, request)

	actualResult := writer.Code

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}
