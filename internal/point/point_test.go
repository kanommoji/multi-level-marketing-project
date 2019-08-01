package point

import (
	"multi-level-marketing-project/database/db"
	"multi-level-marketing-project/models"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_RecordPoint_By_MemberID_99999_Point_50_Should_be_True(t *testing.T) {
	expectedResult := true
	newUserPoint := models.NewUserPoint{
		UserRefferal: 99999,
		NewPoint:     50,
	}
	connection, _ := db.DBConnectionLocal()
	defer connection.Close()
	actualResult := RecordPoint(connection, newUserPoint)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
