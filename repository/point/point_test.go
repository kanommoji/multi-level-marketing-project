package repository

import (
	"multi-level-marketing-project/config"
	"multi-level-marketing-project/database"
	"multi-level-marketing-project/model"
	"testing"
)

func Test_Insert_By_UserReferral_99999_NewPoint_50_Should_Be_True(t *testing.T) {
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

	actualResult := Insert(database, newUserPoint)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
