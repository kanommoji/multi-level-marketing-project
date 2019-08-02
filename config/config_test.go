package config

import (
	"testing"
)

func Test_GetURI_Should_Be_(t *testing.T) {
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
