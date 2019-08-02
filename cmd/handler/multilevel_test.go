package handler

import (
	"bytes"
	"encoding/json"
	"multi-level-marketing-project/model"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Test_GetMember_By_MemberID_10029_Should_Get_Member(t *testing.T) {
	var expectedResult, actualResult model.Member
	response := `{"member_id":10029,"member_name":"ชนา","leader_id":20029,"level":6,"my_point":1000,"monthly_point":0,"team_point":20000,"TeamMember":{"HigherPearl":2,"HigherEmerald":2,"HigherRuby":0}}`
	decodeExpectedResult := json.NewDecoder(bytes.NewReader([]byte(response)))
	_ = decodeExpectedResult.Decode(&expectedResult)
	request := httptest.NewRequest("GET", "/members/10029", nil)
	request = mux.SetURLVars(request, map[string]string{
		"memberID": "10029",
	})
	request.Header.Set("Content-type", "application/json")
	writer := httptest.NewRecorder()
	GetMember(writer, request)

	decodeActualResult := json.NewDecoder(writer.Body)
	_ = decodeActualResult.Decode(&actualResult)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}
