package repository

import (
	"database/sql"
	"multi-level-marketing-project/config"
	"multi-level-marketing-project/model"
)

func Insert(database *sql.DB, newUserPoint model.NewUserPoint) bool {
	statement, err := database.Prepare(`INSERT INTO point_records (member_id, point, create_date) VALUES ( ? , ? , ?)`)
	if err != nil {
		return false
	}
	_, err = statement.Exec(newUserPoint.UserReferral, newUserPoint.NewPoint, config.TimeNow())
	if err != nil {
		return false
	}
	return true
}
