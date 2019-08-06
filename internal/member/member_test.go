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
		Username: "mlm_dev",
		Password: "mlm_dev",
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

func Test_CheckCondition_By_Member_Level_PearlPup_MyPoint_620_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := model.Member{
		Level:   1,
		MyPoint: 620,
	}

	actualResult := CheckCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Member_Level_PearlPup_MyPoint_500_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := model.Member{
		Level:   1,
		MyPoint: 500,
	}

	actualResult := CheckCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Member_Level_PearlJuvenile_MyPoint_1020_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := model.Member{
		Level:   2,
		MyPoint: 1020,
	}

	actualResult := CheckCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Member_Level_PearlJuvenile_MyPoint_570_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := model.Member{
		Level:   2,
		MyPoint: 570,
	}

	actualResult := CheckCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_PearlAlpha_MonthlyPoint_120_TeamPoint_4020_TeamMember_HigherPearl_2_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := model.Member{
		Level:        3,
		TeamPoint:    4020,
		MonthlyPoint: 120,
		TeamMember: model.TeamMember{
			HigherPearl: 2,
		},
	}

	actualResult := CheckCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_PearlAlpha_MonthlyPoint_100_TeamPoint_3550_TeamMember_HigherPearl_2_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := model.Member{
		Level:        3,
		TeamPoint:    3550,
		MonthlyPoint: 100,
		TeamMember: model.TeamMember{
			HigherPearl: 2,
		},
	}

	actualResult := CheckCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_EmeraldAlpha_MonthlyPoint_350_TeamPoint_20000_TeamMember_HigherEmerald_2_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := model.Member{
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

func Test_CheckCondition_By_Level_BlueDiamondJuvenile_MonthlyPoint_2050_TeamPoint_200300_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := model.Member{
		Level:        11,
		TeamPoint:    200300,
		MonthlyPoint: 2050,
	}

	actualResult := CheckCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_BlueDiamondJuvenile_MonthlyPoint_2030_TeamPoint_200000_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := model.Member{
		Level:        11,
		TeamPoint:    200000,
		MonthlyPoint: 2030,
	}

	actualResult := CheckCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_VerifyLevel_By_MemberID_30001_Should_Be_True(t *testing.T) {
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

	actualResult := VerifyLevel(database, memberID)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_Promote_By_MemberID_30001_Should_Be_True(t *testing.T) {
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

	actualResult := Promote(database, memberID)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
