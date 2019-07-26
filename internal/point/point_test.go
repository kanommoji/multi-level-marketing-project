package point

import "testing"

func RecordPoint_By_MemberId_10010_Point_50_Should_Be_True(t *testing.T) {
	expectedResult := true
	action := Action{
		UserReferral:   10010,
		NewPoint: 50,
	}

	actualResult := RecordPoint(action)

	if actualResult != expectedResult {
		t.Errorf("Except %v but get %v", expectedResult, actualResult)
	}
}
