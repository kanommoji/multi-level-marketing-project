package repository

import (
	"database/sql"
)

func GetMyPoint(db *sql.DB, memberID int) int {
	statement, err := db.Prepare(`SELECT SUM(point) FROM point_records WHERE member_id = ?`)
	if err != nil {
		return 0
	}
	row := statement.QueryRow(memberID)
	var myPoint int
	row.Scan(&myPoint)
	return myPoint
}
