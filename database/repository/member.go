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

func GetMonthlyPoint(db *sql.DB, memberID int, month int, year int) int {
	statement, err := db.Prepare(`SELECT SUM(point) FROM mlm.point_records WHERE member_id = ? AND MONTH(create_date) = ? AND YEAR(create_date) = ?`)
	if err != nil {
		return 0
	}
	row := statement.QueryRow(memberID, month, year)
	var monthlyPoint int
	row.Scan(&monthlyPoint)
	return monthlyPoint
}
