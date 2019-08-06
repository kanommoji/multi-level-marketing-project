package config

import (
	"testing"
	"time"
)

func Test_GetURI_Should_Be_Username_root_Password_root_Host_127_0_0_1_Port_3306_Database_mlm(t *testing.T) {
	expectedResult := "root:root@tcp(127.0.0.1:3306)/mlm"
	config := Config{
		Username: "root",
		Password: "root",
		Host:     "127.0.0.1",
		Database: "mlm",
		Port:     "3306",
	}

	actualResult := config.GetURI()

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_TimeNow_Shout_Be_Date_Year_2019_Month_7_Day_2(t *testing.T) {
	expectedResult := time.Date(2019, 7, 2, 0, 0, 0, 0, time.UTC)

	actualResult := TimeNow()

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
