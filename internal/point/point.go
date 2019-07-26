package point

import (
	mlmDB "multi-level-marketing-project/database/db"
	"time"
)

type Action struct {
	UserReferral int `json:"user_referral"`
	NewPoint     int `json:"new_point"`
}

type Record struct {
	Id         int
	MemberId   int
	Point      int
	CreateDate time.Time
}

func RecordPoint(action Action) bool {
	db := mlmDB.Connect()
	defer db.Close()
	statement, _ := db.Prepare(`INSERT INTO mlm.point_records (member_id, point) VALUES (?,?)`)
	_, err := statement.Exec(action.UserReferral, action.NewPoint)
	if err != nil {
		return false
	}
	return true
}
