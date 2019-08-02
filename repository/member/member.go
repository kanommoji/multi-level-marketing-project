package repository

import (
	"database/sql"
	"multi-level-marketing-project/model"
)

func FindMemberByID(database *sql.DB, memberID int) model.Member {
	var member model.Member
	statement, err := database.Prepare(`SELECT * FROM members WHERE id = ?`)
	if err != nil {
		return member
	}
	row := statement.QueryRow(memberID)
	row.Scan(
		&member.MemberID,
		&member.MemberName,
		&member.LeaderID,
		&member.Level,
	)
	return member
}

func GetTeamPoint(database *sql.DB, memberID int) int {
	var teamPoint int
	statement, err := database.Prepare(`SELECT SUM(point) FROM point_records INNER JOIN members ON members.id = point_records.member_id WHERE leader_id = ? OR member_id = ?`)
	if err != nil {
		return teamPoint
	}
	row := statement.QueryRow(memberID, memberID)
	row.Scan(&teamPoint)
	return teamPoint
}
