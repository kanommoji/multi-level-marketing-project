package members

import "testing"

func Test_CheckCondition_By_Member_Level_1_MyPoint_620_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := Member{
		Level:        1,
		MyPoint: 620,
	}

	actualResult := checkCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Member_Level_1_MyPoint_570_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := Member{
		Level:        1,
		MyPoint: 570,
	}

	actualResult := checkCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Member_Level_6_MonthlyPoint_400_TeamPoint_20050_TeamMemberHigherEmerald_2_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := Member{
		Level:        6,
		MonthlyPoint: 400,
		TeamPoint:    20050,
		TeamMember: TeamMember{
			HigherEmerald: 2,
		},
	}

	actualResult := checkCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Member_Level_6_MonthlyPoint_400_TeamPoint_19050_TeamMemberHigherEmerald_2_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := Member{
		Level:        6,
		MonthlyPoint: 400,
		TeamPoint:    19050,
		TeamMember: TeamMember{
			HigherEmerald: 2,
		},
	}

	actualResult := checkCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CountTeamMember_By_Member_Id_10029_Should_Be_TeamMember_HigherPearl_2(t *testing.T) {
	expectedResult := TeamMember{
		HigherPearl:   2,
		HigherEmerald: 2,
	}

	actualResult := CountTeamMember(10029)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
func Test_GetMonthlyPoint_By_Member_Id_10029_Month_7_Year_2019_Should_Be_350(t *testing.T) {
	expectedResult := 350

	actualResult := getMonthlyPoint(10029, 7, 2019)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetMonthlyPoint_By_Member_Id_10009_Month_7_Year_2019_Should_Be_100(t *testing.T) {
	expectedResult := 100

	actualResult := getMonthlyPoint(10009, 7, 2019)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetMyPoint_By_Member_Id_10029_Should_Be_1000(t *testing.T) {
	expectedResult := 1000

	actualResult := getMyPoint(10029)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_Promote_By_Member_Id_10029_Level_6_Should_Be_True(t *testing.T) {
	expectedResult := true

	actualResult := Promote(Member{MemberId: 10029, Level: 6})

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}

	Demote(Member{MemberId: 10029, Level: 7})
}

func Test_FindMemberAlpha_By_Member_Id_10029_Should_Be_Member(t *testing.T) {
	expectedResult := Member{
		MemberId:     10029,
		MemberName:   "ชนา",
		LeaderId:     20029,
		Level:        6,
		MyPoint:      1000,
		MonthlyPoint: 350,
		TeamPoint:    20000,
		TeamMember: TeamMember{
			HigherPearl:   2,
			HigherEmerald: 2,
		},
	}

	actualResult := FindMember(10029)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetTeamPoint_By_Member_Id_10029_Should_Be_20000(t *testing.T) {
	expectedResult := 20000

	actualResult := getTeamPoint(10029)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_VerifyLevel_By_Member_Id_30001_Should_Be_True(t *testing.T) {
	expectedResult := true

	actualResult := VerifyLevel(30001)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}

	Demote(Member{MemberId: 30001, Level: 7})
}
