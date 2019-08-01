package repository

import (
	"database/sql"
	"multi-level-marketing-project/models"
)

func Insert(db *sql.DB, newUserPoint models.NewUserPoint) bool {
	statements, err := db.Prepare(`INSERT INTO point_records (member_id, point) VALUES (?,?)`)
	if err != nil {
		return false
	}
	_, err = statements.Exec(newUserPoint.UserRefferal, newUserPoint.NewPoint)
	if err != nil {
		return false
	}
	return true
}
