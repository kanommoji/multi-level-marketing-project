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

func Test_CheckCondition_By_Level_EmeraldPup_MonthlyPoint_160_TeamPoint_8030_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := model.Member{
		Level:        4,
		TeamPoint:    8030,
		MonthlyPoint: 160,
	}

	actualResult := CheckCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_EmeraldPup_MonthlyPoint_150_TeamPoint_7520_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := model.Member{
		Level:        4,
		TeamPoint:    7520,
		MonthlyPoint: 150,
	}

	actualResult := CheckCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_EmeraldJuvenile_MonthlyPoint_200_TeamPoint_12050_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := model.Member{
		Level:        5,
		TeamPoint:    12050,
		MonthlyPoint: 200,
	}

	actualResult := CheckCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_EmeraldJuvenile_MonthlyPoint_200_TeamPoint_10050_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := model.Member{
		Level:        5,
		TeamPoint:    10050,
		MonthlyPoint: 200,
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

func Test_CheckCondition_By_Level_RubyPup_MonthlyPoint_600_TeamPoint_40050_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := model.Member{
		Level:        7,
		TeamPoint:    40050,
		MonthlyPoint: 600,
	}

	actualResult := CheckCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_RubyPup_MonthlyPoint_600_TeamPoint_30050_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := model.Member{
		Level:        7,
		TeamPoint:    30050,
		MonthlyPoint: 600,
	}

	actualResult := CheckCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_RubyJuvenile_MonthlyPoint_800_TeamPoint_100050_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := model.Member{
		Level:        8,
		TeamPoint:    100050,
		MonthlyPoint: 800,
	}

	actualResult := CheckCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_RubyJuvenile_MonthlyPoint_800_TeamPoint_60050_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := model.Member{
		Level:        8,
		TeamPoint:    60050,
		MonthlyPoint: 800,
	}

	actualResult := CheckCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_RubyAlpha_MonthlyPoint_1000_TeamPoint_100050_TeamMember_HigherRuby_2_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := model.Member{
		Level:        9,
		TeamPoint:    100050,
		MonthlyPoint: 1000,
		TeamMember: model.TeamMember{
			HigherRuby: 2,
		},
	}

	actualResult := CheckCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_RubyAlpha_MonthlyPoint_970_TeamPoint_9020_TeamMember_HigherRuby_1_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := model.Member{
		Level:        9,
		TeamPoint:    9020,
		MonthlyPoint: 970,
		TeamMember: model.TeamMember{
			HigherRuby: 1,
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

func Test_CheckCondition_By_Level_BlueDiamondPup_MonthlyPoint_1600_TeamPoint_160050_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := model.Member{
		Level:        10,
		TeamPoint:    160050,
		MonthlyPoint: 1600,
		TeamMember: model.TeamMember{
			HigherRuby: 2,
		},
	}

	actualResult := CheckCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_BlueDiamondPup_MonthlyPoint_1600_TeamPoint_160050_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := model.Member{
		Level:        10,
		TeamPoint:    150050,
		MonthlyPoint: 1600,
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

func Test_CheckConditionDemote_By_Level_EmeraldPup_MonthlyPoint_70_TeamMemberHigherPearl_2_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := model.Member{
		Level:        4,
		MonthlyPoint: 70,
		TeamMember: model.TeamMember{
			HigherPearl: 2,
		},
	}

	actualResult := CheckConditionDemote(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckConditionDemote_By_Level_EmeraldPup_MonthlyPoint_100_TeamMemberHigherPearl_1_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := model.Member{
		Level:        4,
		MonthlyPoint: 100,
		TeamMember: model.TeamMember{
			HigherPearl: 1,
		},
	}

	actualResult := CheckConditionDemote(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckConditionDemote_By_Level_EmeraldPup_MonthlyPoint_150_TeamMemberHigherPearl_2_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := model.Member{
		Level:        4,
		MonthlyPoint: 150,
		TeamMember: model.TeamMember{
			HigherPearl: 2,
		},
	}

	actualResult := CheckConditionDemote(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_Demote_By_MemberID_29001_Should_Be_True(t *testing.T) {
	expectedResult := true
	memberID := 29001
	config := config.Config{
		Username: "mlm_dev",
		Password: "mlm_dev",
		Host:     "127.0.0.1",
		Database: "mlm",
		Port:     "3306",
	}
	database, _ := database.DBConnect(config.GetURI())

	actualResult := Demote(database, memberID)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_VerifyLevelDemote_By_MemberID_40001_Should_Be_True(t *testing.T) {
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

	actualResult := VerifyLevelDemote(database, memberID)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
