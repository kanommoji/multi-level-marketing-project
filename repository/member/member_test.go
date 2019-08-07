package repository

import (
	"multi-level-marketing-project/config"
	"multi-level-marketing-project/database"
	"multi-level-marketing-project/model"
	"testing"
)

func Test_FindMemberByID_By_MemberID_10029_Should_Be_MemberName_ชนา_LeaderID_20029_Level_6(t *testing.T) {
	expectedResult := model.Member{
		MemberID:   10029,
		MemberName: "ชนา",
		LeaderID:   20029,
		Level:      6,
	}
	memberID := 10029
	config := config.Config{
		Username: "mlm_dev",
		Password: "mlm_dev",
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

func Test_GetMyPoint_By_MemberID_10029_Should_Be_1000(t *testing.T) {
	expectedResult := 1000
	memberID := 10029
	config := config.Config{
		Username: "mlm_dev",
		Password: "mlm_dev",
		Host:     "127.0.0.1",
		Database: "mlm",
		Port:     "3306",
	}
	database, _ := database.DBConnect(config.GetURI())

	actualResult := GetMyPoint(database, memberID)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetTeamPoint_By_MemberID_10029_Should_Be_20000(t *testing.T) {
	expectedResult := 20000
	memberID := 10029
	config := config.Config{
		Username: "mlm_dev",
		Password: "mlm_dev",
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

func Test_GetMonthlyPoint_By_MemberID_10029_Month_7_Year_2019_Should_Be_350(t *testing.T) {
	expectedResult := 350
	memberID := 10029
	month := 7
	year := 2019
	config := config.Config{
		Username: "mlm_dev",
		Password: "mlm_dev",
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

func Test_CountTeamMember_By_MemberID_10029_Should_Be_TeamMember_HigherPearl_2_HigherEmerald_2(t *testing.T) {
	expectedResult := model.TeamMember{
		HigherPearl:   2,
		HigherEmerald: 2,
	}
	memberID := 10029
	config := config.Config{
		Username: "mlm_dev",
		Password: "mlm_dev",
		Host:     "127.0.0.1",
		Database: "mlm",
		Port:     "3306",
	}
	database, _ := database.DBConnect(config.GetURI())

	actualResult := CountTeamMember(database, memberID)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_UpdateLevelPlusOne_By_MemberID_30001_Should_Be_True(t *testing.T) {
	expectedResult := true
	memberID := 30001
	config := config.Config{
		Username: "mlm_dev",
		Password: "mlm_dev",
		Host:     "127.0.0.1",
		Database: "mlm",
		Port:     "3306",
	}
	database, _ := database.DBConnect(config.GetURI())

	actualResult := UpdateLevelPlusOne(database, memberID)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_UpdateLevelDownOne_By_MemberID_40001_Should_Be_True(t *testing.T) {
	expectedResult := true
	memberID := 40001
	config := config.Config{
		Username: "mlm_dev",
		Password: "mlm_dev",
		Host:     "127.0.0.1",
		Database: "mlm",
		Port:     "3306",
	}
	database, _ := database.DBConnect(config.GetURI())

	actualResult := UpdateLevelDownOne(database, memberID)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
