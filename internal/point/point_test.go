package point

import (
	"database/sql"
	"multi-level-marketing-project/internal/members"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql",
		"root:1640@tcp(127.0.0.1:3306)/mlm")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Test_RecordPoint_By_MemberID_99999_Point_50_Should_be_True(t *testing.T) {
	expectedResult := true
	newUserPoint := members.NewUserPoint{
		UserRefferal: 99999,
		NewPoint:     50,
	}
	db, _ := DBConnection()
	defer db.Close()
	actualResult := RecordPoint(db, newUserPoint)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
