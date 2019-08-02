package repository

import (
	"multi-level-marketing-project/config"
	"multi-level-marketing-project/database"
	"multi-level-marketing-project/model"
	"testing"
)

func Test_FindMemberByID_By_Member_10029_Should_Be_MemberName_ชนา_LeaderID_20029_Level_6(t *testing.T) {
	expectedResult := model.Member{
		MemberID:   10029,
		MemberName: "ชนา",
		LeaderID:   20029,
		Level:      6,
	}
	memberID := 10029
	config := config.Config{
		Username: "root",
		Password: "root",
		Host:     "127.0.0.1",
		Database: "mlm",
		Port:     "3306",
	}
	database, _ := database.DBConnect(config.GetURI())

	actualResult := FindMemberByID(database, memberID)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetMyPoint_By_Member_Id_10029_Should_Be_1000(t *testing.T) {
	expectedResult := 1000
	memberID := 10029
	config := config.Config{
		Username: "root",
		Password: "root",
		Host:     "127.0.0.1",
		Database: "mlm",
		Port:     "3306",
	}
	database, _ := database.DBConnect(config.GetURI())

	actualResult := getMyPoint(database, memberID)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetTeamPoint_By_Member_10029_Should_Be_TeamPoint_20000(t *testing.T) {
	expectedResult := 20000
	memberID := 10029
	config := config.Config{
		Username: "root",
		Password: "root",
		Host:     "127.0.0.1",
		Database: "mlm",
		Port:     "3306",
	}
	database, _ := database.DBConnect(config.GetURI())

	actualResult := GetTeamPoint(database, memberID)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetMonthlyPoint_By_Member_10029_Month_7_Year_2019_Should_Be_MonthlyPoint_350(t *testing.T) {
	expectedResult := 350
	memberID := 10029
	month := 7
	year := 2019
	config := config.Config{
		Username: "root",
		Password: "root",
		Host:     "127.0.0.1",
		Database: "mlm",
		Port:     "3306",
	}
	database, _ := database.DBConnect(config.GetURI())

	actualResult := GetMonthlyPoint(database, memberID, month, year)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
