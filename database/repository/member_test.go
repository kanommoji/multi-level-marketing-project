package repository

import (
	connection "multi-level-marketing-project/database/db"
	"multi-level-marketing-project/models"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_GetMyPoint_By_MemberID_10029_Should_be_Point_1000(t *testing.T) {
	expectedResult := 1000

	db, _ := connection.DBConnectionLocal()
	defer db.Close()
	actualResult := GetMyPoint(db, 10029)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetMonthlyPoint_By_MemberID_10029_Should_be_MonthlyPoint_350(t *testing.T) {
	expectedResult := 350

	db, _ := connection.DBConnectionLocal()
	defer db.Close()
	actualResult := GetMonthlyPoint(db, 10029, 7, 2019)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CountTeamMember_By_MemberID_10029_Should_be_TeamMember_HigherPeal_2_HigherEmerald_2(t *testing.T) {
	expectedResult := models.TeamMember{
		HigherPearl:   2,
		HigherEmerald: 2,
	}

	db, _ := connection.DBConnectionLocal()
	defer db.Close()
	actualResult := CountTeamMember(db, 10029)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetTeamPoint_By_MemberID_10029_Should_be_TeamPoint_20000(t *testing.T) {
	expectedResult := 20000

	db, _ := connection.DBConnectionLocal()
	defer db.Close()
	actualResult := GetTeamPoint(db, 10029)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetMemberData_By_MemberID_10029_Should_be_Member(t *testing.T) {
	expectedResult := models.Member{
		MemberID:   10029,
		MemberName: "ชนา",
		LeaderID:   20029,
		Level:      6,
	}

	db, _ := connection.DBConnectionLocal()
	defer db.Close()
	actualResult := GetMemberData(db, 10029)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_Promote_By_MemberID_10029_Level_6_Should_be_True(t *testing.T) {
	expectedResult := true

	db, _ := connection.DBConnectionLocal()
	defer db.Close()

	actualResult := Promote(db, models.Member{MemberID: 10029, Level: 6})

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}

	Demote(db, models.Member{MemberID: 10029, Level: 7})
}
