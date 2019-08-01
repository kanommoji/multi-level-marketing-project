package repository

import (
	connection "multi-level-marketing-project/database/db"
	"multi-level-marketing-project/models"
	"testing"
)

func Test_Insert_By_MemberID_99999_Point_50_Should_be_True(t *testing.T) {
	expectedResult := true
	newUserPoint := models.NewUserPoint{
		UserRefferal: 99999,
		NewPoint:     50,
	}
	db, _ := connection.DBConnection()
	defer db.Close()
	actualResult := Insert(db, newUserPoint)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
