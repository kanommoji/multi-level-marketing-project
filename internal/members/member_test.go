package members

import "testing"

func Test_CheckCondition_By_Level_1_MyPoint_620_Should_be_True(t *testing.T) {
	expectedResult := true
	member := Member{
		Level:   1,
		MyPoint: 620,
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_1_MyPoint_600_Should_be_False(t *testing.T) {
	expectedResult := false
	member := Member{
		Level:   1,
		MyPoint: 600,
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_1_MyPoint_40_Should_be_False(t *testing.T) {
	expectedResult := false
	member := Member{
		Level:   1,
		MyPoint: 40,
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_3_MonthlyPoint_120_TeamPoint_4020_TeamMember_HigherPearl_2_Should_be_True(t *testing.T) {
	expectedResult := true
	member := Member{
		Level:        3,
		MonthlyPoint: 120,
		TeamPoint:    4020,
		TeamMember: TeamMember{
			HigherPearl: 2,
		},
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_3_MonthlyPoint_100_TeamPoint_3550_TeamMember_HigherPearl_1_Should_be_False(t *testing.T) {
	expectedResult := false
	member := Member{
		Level:        3,
		MonthlyPoint: 100,
		TeamPoint:    3550,
		TeamMember: TeamMember{
			HigherPearl: 1,
		},
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_3_MonthlyPoint_140_TeamPoint_5520_TeamMember_HigherPearl_1_Should_be_False(t *testing.T) {
	expectedResult := false
	member := Member{
		Level:        3,
		MonthlyPoint: 140,
		TeamPoint:    5520,
		TeamMember: TeamMember{
			HigherPearl: 1,
		},
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
func Test_CheckCondition_By_Level_6_MonthlyPoint_400_TeamPoint_20050_TeamMember_HigherEmerald_2_Should_be_True(t *testing.T) {
	expectedResult := true
	member := Member{
		Level:        6,
		MonthlyPoint: 400,
		TeamPoint:    20050,
		TeamMember: TeamMember{
			HigherEmerald: 2,
		},
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_6_MonthlyPoint_400_TeamPoint_20000_TeamMember_HigherEmerald_2_Should_be_False(t *testing.T) {
	expectedResult := false
	member := Member{
		Level:        6,
		MonthlyPoint: 400,
		TeamPoint:    20000,
		TeamMember: TeamMember{
			HigherEmerald: 2,
		},
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_6_MonthlyPoint_370_TeamPoint_15020_TeamMember_HigherEmerald_2_Should_be_False(t *testing.T) {
	expectedResult := false
	member := Member{
		Level:        6,
		MonthlyPoint: 370,
		TeamPoint:    15020,
		TeamMember: TeamMember{
			HigherEmerald: 2,
		},
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_11_MonthlyPoint_2050_TeamPoint_200300_Should_be_True(t *testing.T) {
	expectedResult := true
	member := Member{
		Level:        11,
		MonthlyPoint: 2050,
		TeamPoint:    200300,
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_11_MonthlyPoint_2000_TeamPoint_200000_Should_be_False(t *testing.T) {
	expectedResult := false
	member := Member{
		Level:        11,
		MonthlyPoint: 2000,
		TeamPoint:    200000,
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CheckCondition_By_Level_11_MonthlyPoint_2050_TeamPoint_200050_Should_be_True(t *testing.T) {
	expectedResult := false
	member := Member{
		Level:        11,
		MonthlyPoint: 2050,
		TeamPoint:    150050,
	}

	actualResult := CheckCondition(member)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
