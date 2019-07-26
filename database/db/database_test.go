package db

import "testing"

func Test_OpenDataBase_Then_Can_Sent_Query_To_DB_Without_Error(t *testing.T) {
	db := Connect()
	defer db.Close()
	_, err := db.Query("SELECT * FROM members")

	if err != nil {
		t.Errorf("Can't connect to database , %v", err)
	}
}
