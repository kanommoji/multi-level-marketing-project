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
