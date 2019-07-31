package members

import "testing"

func Test_CheckCondition_By_Member_Level_1_MyPoint_620_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := Member{
		Level:   1,
		MyPoint: 620,
	}

	actualResult := checkCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Member_Level_1_MyPoint_600_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := Member{
		Level:   1,
		MyPoint: 600,
	}

	actualResult := checkCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Member_Level_1_MyPoint_570_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := Member{
		Level:   1,
		MyPoint: 570,
	}

	actualResult := checkCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Member_Level_1_MyPoint_40_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := Member{
		Level:   1,
		MyPoint: 40,
	}

	actualResult := checkCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Member_Level_3_MonthlyPoint_120_TeamPoint_4020_TeamMemberHigherPearl_2_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := Member{
		Level:        3,
		MonthlyPoint: 120,
		TeamPoint:    4020,
		TeamMember: TeamMember{
			HigherPearl: 2,
		},
	}

	actualResult := checkCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Member_Level_3_MonthlyPoint_140_TeamPoint_5520_TeamMemberHigherPearl_2_Should_Be_True(t *testing.T) {
	expectedResult := true
	member := Member{
		Level:        3,
		MonthlyPoint: 140,
		TeamPoint:    5520,
		TeamMember: TeamMember{
			HigherPearl: 2,
		},
	}

	actualResult := checkCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Member_Level_3_MonthlyPoint_100_TeamPoint_3550_TeamMemberHigherPearl_1_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := Member{
		Level:        3,
		MonthlyPoint: 100,
		TeamPoint:    3550,
		TeamMember: TeamMember{
			HigherPearl: 1,
		},
	}

	actualResult := checkCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Member_Level_3_MonthlyPoint_120_TeamPoint_4020_TeamMemberHigherPearl_1_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := Member{
		Level:        3,
		MonthlyPoint: 120,
		TeamPoint:    4020,
		TeamMember: TeamMember{
			HigherPearl: 1,
		},
	}

	actualResult := checkCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Member_Level_3_MonthlyPoint_140_TeamPoint_5520_TeamMemberHigherPearl_1_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := Member{
		Level:        3,
		MonthlyPoint: 140,
		TeamPoint:    5520,
		TeamMember: TeamMember{
			HigherPearl: 1,
		},
	}

	actualResult := checkCondition(member)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Member_Level_3_MonthlyPoint_70_TeamPoint_3520_TeamMemberHigherPearl_2_Should_Be_False(t *testing.T) {
	expectedResult := false
	member := Member{
		Level:        3,
		MonthlyPoint: 70,
		TeamPoint:    3520,
		TeamMember: TeamMember{
			HigherPearl: 2,
		},
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

func Test_CountTeamMember_By_Member_Id_10029_Should_Be_TeamMember_HigherPearl_2_HigherEmerald_2(t *testing.T) {
	expectedResult := TeamMember{
		HigherPearl:   2,
		HigherEmerald: 2,
	}

	actualResult := countTeamMember(10029)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetMonthlyPoint_By_Member_Id_10029_Month_7_Year_2019_Should_Be_MonthlyPoint_350(t *testing.T) {
	expectedResult := 350

	actualResult := getMonthlyPoint(10029, 7, 2019)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetMonthlyPoint_By_Member_Id_10009_Month_7_Year_2019_Should_Be_MonthlyPoint_100(t *testing.T) {
	expectedResult := 100

	actualResult := getMonthlyPoint(10009, 7, 2019)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetMyPoint_By_Member_Id_10029_Should_Be_MyPoint_1000(t *testing.T) {
	expectedResult := 1000

	actualResult := getMyPoint(10029)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_Promote_By_Member_Id_10029_Level_6_Should_Be_Level_7(t *testing.T) {
	expectedResult := 7

	promote(Member{MemberId: 10029, Level: 6})
	actualResult := FindMember(10029).Level

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}

	demote(Member{MemberId: 10029, Level: 7})
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

func Test_GetTeamPoint_By_Member_Id_10029_Should_Be_TeamPoint_20000(t *testing.T) {
	expectedResult := 20000

	actualResult := getTeamPoint(10029)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_VerifyLevel_By_Member_Id_30001_Should_Be_Level_7(t *testing.T) {
	expectedResult := 7

	VerifyLevel(30001)
	actualResult := FindMember(30001).Level

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}

	demote(Member{MemberId: 30001, Level: 7})
}
