package repository

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_GetMyPoint_By_MemberID_10029_Should_be_Point_1000(t *testing.T) {
	expectedResult := 1000

	db, _ := DBConnection()
	defer db.Close()
	actualResult := GetMyPoint(db, 10029)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
