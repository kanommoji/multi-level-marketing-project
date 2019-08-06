package point

import (
	"multi-level-marketing-project/config"
	"multi-level-marketing-project/database"
	"multi-level-marketing-project/model"
	"testing"
)

func Test_RecordPoint_By_MemberID_99999_Point_50_Should_Be_True(t *testing.T) {
	expectedResult := true
	config := config.Config{
		Username: "mlm_dev",
		Password: "mlm_dev",
		Host:     "127.0.0.1",
		Database: "mlm",
		Port:     "3306",
	}
	database, _ := database.DBConnect(config.GetURI())
	newUserPoint := model.NewUserPoint{
		UserReferral: 99999,
		NewPoint:     50,
	}

	actualResult := RecordPoint(database, newUserPoint)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_RecordPoint_By_MemberID_00000_Point_50_Should_Be_False(t *testing.T) {
	expectedResult := false
	config := config.Config{
		Username: "mlm_dev",
		Password: "mlm_dev",
		Host:     "127.0.0.1",
		Database: "mlm",
		Port:     "3306",
	}
	database, _ := database.DBConnect(config.GetURI())
	newUserPoint := model.NewUserPoint{
		UserReferral: 00000,
		NewPoint:     50,
	}

	actualResult := RecordPoint(database, newUserPoint)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
