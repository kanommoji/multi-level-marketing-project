package point

import "testing"

func Test_RecordPoint_By_Member_Id_99999_Point_50_Should_Be_True(t *testing.T) {
	expectedResult := true
	action := Action{
		UserReferral: 99999,
		NewPoint:     50,
	}

	actualResult := RecordPoint(action)

	if actualResult != expectedResult {
		t.Errorf("Except %v but get %v", expectedResult, actualResult)
	}
}
