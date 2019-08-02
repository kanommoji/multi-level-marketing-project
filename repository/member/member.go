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

func getMyPoint(database *sql.DB, memberID int) int {
	var myPoint int
	statement, err := database.Prepare(`SELECT SUM(point) AS my_point FROM point_records WHERE member_id = ?`)
	if err != nil {
		return 0
	}
	row := statement.QueryRow(memberID)
	row.Scan(&myPoint)
	if err != nil {
		return 0
	}
	return myPoint
}

func GetTeamPoint(database *sql.DB, memberID int) int {
	var teamPoint int
	statement, err := database.Prepare(`SELECT SUM(point) FROM point_records INNER JOIN members ON members.id = point_records.member_id WHERE leader_id = ? OR member_id = ?`)
	if err != nil {
		return 0
	}
	row := statement.QueryRow(memberID, memberID)
	row.Scan(&teamPoint)
	return teamPoint
}

func GetMonthlyPoint(database *sql.DB, memberID int, month int, year int) int {
	var monthlyPoint int
	statement, err := database.Prepare(`SELECT SUM(point) FROM point_records WHERE member_id = ? AND MONTH(create_date)= ? AND YEAR(create_date)= ?`)
	if err != nil {
		return 0
	}
	row := statement.QueryRow(memberID, month, year)
	row.Scan(&monthlyPoint)
	return monthlyPoint
}

func CountTeamMember(database *sql.DB, memberID int) model.TeamMember {
	var teamMember model.TeamMember
	statement, err := database.Prepare(`SELECT COUNT(id) FROM members WHERE leader_id = ? AND level >= ?`)
	if err != nil {
		return teamMember
	}
	rowTeamMemberHigherPearl := statement.QueryRow(memberID, 1)
	rowTeamMemberHigherPearl.Scan(&teamMember.HigherPearl)

	rowTeamMemberHigherEmerald := statement.QueryRow(memberID, 4)
	rowTeamMemberHigherEmerald.Scan(&teamMember.HigherEmerald)

	rowTeamMemberHigherRuby := statement.QueryRow(memberID, 7)
	rowTeamMemberHigherRuby.Scan(&teamMember.HigherRuby)

	return teamMember
}
