package members

import (
	connection "multi-level-marketing-project/database/db"
	"multi-level-marketing-project/database/repository"
	"multi-level-marketing-project/models"
	"testing"
)

func Test_CheckCondition_By_Level_PearlPup_MyPoint_620_Should_be_True(t *testing.T) {
	expectedResult := true
	member := models.Member{
		Level:   1,
		MyPoint: 620,
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_PearlPup_MyPoint_600_Should_be_False(t *testing.T) {
	expectedResult := false
	member := models.Member{
		Level:   1,
		MyPoint: 600,
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_PearlPup_MyPoint_40_Should_be_False(t *testing.T) {
	expectedResult := false
	member := models.Member{
		Level:   1,
		MyPoint: 40,
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_PearlAlpha_MonthlyPoint_120_TeamPoint_4020_TeamMember_HigherPearl_2_Should_be_True(t *testing.T) {
	expectedResult := true
	member := models.Member{
		Level:        3,
		MonthlyPoint: 120,
		TeamPoint:    4020,
		TeamMember: models.TeamMember{
			HigherPearl: 2,
		},
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_PearlAlpha_MonthlyPoint_100_TeamPoint_3550_TeamMember_HigherPearl_1_Should_be_False(t *testing.T) {
	expectedResult := false
	member := models.Member{
		Level:        3,
		MonthlyPoint: 100,
		TeamPoint:    3550,
		TeamMember: models.TeamMember{
			HigherPearl: 1,
		},
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_PearlAlpha_MonthlyPoint_140_TeamPoint_5520_TeamMember_HigherPearl_1_Should_be_False(t *testing.T) {
	expectedResult := false
	member := models.Member{
		Level:        3,
		MonthlyPoint: 140,
		TeamPoint:    5520,
		TeamMember: models.TeamMember{
			HigherPearl: 1,
		},
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
func Test_CheckCondition_By_Level_EmeraldAlpha_MonthlyPoint_400_TeamPoint_20050_TeamMember_HigherEmerald_2_Should_be_True(t *testing.T) {
	expectedResult := true
	member := models.Member{
		Level:        6,
		MonthlyPoint: 400,
		TeamPoint:    20050,
		TeamMember: models.TeamMember{
			HigherEmerald: 2,
		},
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_EmeraldAlpha_MonthlyPoint_400_TeamPoint_20000_TeamMember_HigherEmerald_2_Should_be_False(t *testing.T) {
	expectedResult := false
	member := models.Member{
		Level:        6,
		MonthlyPoint: 400,
		TeamPoint:    20000,
		TeamMember: models.TeamMember{
			HigherEmerald: 2,
		},
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_EmeraldAlpha_MonthlyPoint_370_TeamPoint_15020_TeamMember_HigherEmerald_2_Should_be_False(t *testing.T) {
	expectedResult := false
	member := models.Member{
		Level:        6,
		MonthlyPoint: 370,
		TeamPoint:    15020,
		TeamMember: models.TeamMember{
			HigherEmerald: 2,
		},
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_BlueDiamondJuvenile_MonthlyPoint_2050_TeamPoint_200300_Should_be_True(t *testing.T) {
	expectedResult := true
	member := models.Member{
		Level:        11,
		MonthlyPoint: 2050,
		TeamPoint:    200300,
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_BlueDiamondJuvenile_MonthlyPoint_2000_TeamPoint_200000_Should_be_False(t *testing.T) {
	expectedResult := false
	member := models.Member{
		Level:        11,
		MonthlyPoint: 2000,
		TeamPoint:    200000,
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_BlueDiamondJuvenile_MonthlyPoint_2050_TeamPoint_200050_Should_be_True(t *testing.T) {
	expectedResult := false
	member := models.Member{
		Level:        11,
		MonthlyPoint: 2050,
		TeamPoint:    150050,
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_FindMember_By_MemberID_10029_Should_be_Member(t *testing.T) {
	expectedResult := models.Member{
		MemberID:     10029,
		MemberName:   "ชนา",
		LeaderID:     20029,
		Level:        6,
		MyPoint:      1000,
		MonthlyPoint: 0,
		TeamPoint:    20000,
		TeamMember: models.TeamMember{
			HigherPearl:   2,
			HigherEmerald: 2,
		},
	}

	db, _ := connection.DBConnectionLocal()
	defer db.Close()
	actualResult := FindMember(db, 10029)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_VerifyLevel_By_MemberID_30001_Should_be_True(t *testing.T) {
	expectedResult := true

	db, _ := connection.DBConnectionLocal()
	defer db.Close()
	actualResult := VerifyLevel(db, 30001)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}

	repository.Demote(db, models.Member{MemberID: 30001, Level: 7})
}
