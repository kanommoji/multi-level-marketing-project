package handler

import (
	"net/http/httptest"
	"testing"
)

func Test_GetMember_By_MemberID_10029_Should_Get_Member(t *testing.T) {
	expectedResult := `{member_id:10029,member_name:ชนา,leader_id:20029,level:6,my_point:1000,monthly_point:350,team_point:20000}`
	request := httptest.NewRequest("GET", "/members/10029", nil)
	request.Header.Set("Content-type", "application/json")
	writer := httptest.NewRecorder()
	GetMember(writer, request)

	actualResult := writer.Body.String()

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}
