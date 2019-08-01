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
