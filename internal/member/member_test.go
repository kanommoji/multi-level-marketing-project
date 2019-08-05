package member

import (
	"multi-level-marketing-project/config"
	"multi-level-marketing-project/database"
	"multi-level-marketing-project/model"
	"testing"
)

func Test_FindMember_By_MemberID_10029_Should_Be_MemberName_ชนา_LeaderID_20029_Level_6(t *testing.T) {
	expectedResult := model.Member{
		MemberID:     10029,
		MemberName:   "ชนา",
		LeaderID:     20029,
		Level:        6,
		TeamPoint:    20000,
		MyPoint:      1000,
		MonthlyPoint: 350,
		TeamMember: model.TeamMember{
			HigherPearl:   2,
			HigherEmerald: 2,
		},
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

	actualResult := FindMember(database, memberID)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_MemberID_10029_Level_6_MonthlyPoint_350_Month_7_Year_2019_TeamPoint_20000_TeamMember_HigherEmerald_2_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := model.Member{
		MemberID:     10029,
		Level:        6,
		TeamPoint:    20050,
		MonthlyPoint: 400,
		TeamMember: model.TeamMember{
			HigherEmerald: 2,
		},
	}

	actualResult := CheckCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_VerifyLevel_By_UserReferral_10029_Should_Be_True(t *testing.T) {
	expectedResult := true
	memberID := 10029
	config := config.Config{
		Username: "root",
		Password: "root",
		Host:     "127.0.0.1",
		Database: "mlm",
		Port:     "3306",
	}
	database, _ := database.DBConnect(config.GetURI())

	actualResult := VerifyLevel(database, memberID)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
