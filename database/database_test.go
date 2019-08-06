package database

import "testing"

func Test_DBConnect_Should_Be_Database(t *testing.T) {
	url := "mlm_dev:mlm_dev@tcp(127.0.0.1:3306)/mlm"

	_, err := DBConnect(url)

	if err != nil {
		t.Errorf("database cann't connect get error %v", err.Error())
	}
}
