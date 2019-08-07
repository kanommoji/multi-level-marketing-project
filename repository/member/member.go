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
	err = row.Scan(
		&member.MemberID,
		&member.MemberName,
		&member.LeaderID,
		&member.Level,
	)
	if err != nil {
		return member
	}
	return member
}

func GetMyPoint(database *sql.DB, memberID int) int {
	var myPoint int
	statement, err := database.Prepare(`SELECT SUM(point) AS my_point FROM point_records WHERE member_id = ?`)
	if err != nil {
		return 0
	}
	row := statement.QueryRow(memberID)
	err = row.Scan(&myPoint)
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
	err = row.Scan(&teamPoint)
	if err != nil {
		return 0
	}
	return teamPoint
}

func GetMonthlyPoint(database *sql.DB, memberID int, month int, year int) int {
	var monthlyPoint int
	statement, err := database.Prepare(`SELECT SUM(point) FROM point_records WHERE member_id = ? AND MONTH(create_date)= ? AND YEAR(create_date)= ?`)
	if err != nil {
		return 0
	}
	row := statement.QueryRow(memberID, month, year)
	err = row.Scan(&monthlyPoint)
	if err != nil {
		return 0
	}
	return monthlyPoint
}

func CountTeamMember(database *sql.DB, memberID int) model.TeamMember {
	var teamMember model.TeamMember
	statement, err := database.Prepare(`SELECT COUNT(id) FROM members WHERE leader_id = ? AND level >= ?`)
	if err != nil {
		return teamMember
	}
	rowTeamMemberHigherPearl := statement.QueryRow(memberID, 1)
	err = rowTeamMemberHigherPearl.Scan(&teamMember.HigherPearl)
	if err != nil {
		return teamMember
	}
	rowTeamMemberHigherEmerald := statement.QueryRow(memberID, 4)
	err = rowTeamMemberHigherEmerald.Scan(&teamMember.HigherEmerald)
	if err != nil {
		return teamMember
	}
	rowTeamMemberHigherRuby := statement.QueryRow(memberID, 7)
	err = rowTeamMemberHigherRuby.Scan(&teamMember.HigherRuby)
	if err != nil {
		return teamMember
	}
	return teamMember
}

func UpdateLevelPlusOne(database *sql.DB, memberID int) bool {
	statement, err := database.Prepare(`UPDATE members SET level = level + 1 WHERE id = ?`)
	if err != nil {
		return false
	}
	_, err = statement.Exec(memberID)
	if err != nil {
		return false
	}
	return true
}

func UpdateLevelDownOne(database *sql.DB, memberID int) bool {
	statement, err := database.Prepare(`UPDATE members SET level = level - 1 WHERE id = ?`)
	if err != nil {
		return false
	}
	_, err = statement.Exec(memberID)
	if err != nil {
		return false
	}
	return true
}
