package members

import (
	"log"
	mlmDB "multi-level-marketing-project/database/db"
	"time"
)

const (
	levelPearlPup     = 1
	levelEmeraldPup   = 4
	levelEmeraldAlpha = 6
	levelRubyPup      = 7

	conditionMonthlyPointOfRubyPup            = 400
	conditionTeamPointOfRubyPup               = 20000
	conditionTeamMemberHigherEmeraldOfRubyPup = 2
)

type Member struct {
	MemberId     int    `json:"member_id"`
	MemberName   string `json:"member_name"`
	LeaderId     int    `json:"leader_id"`
	Level        int    `json:"level"`
	MyPoint      int    `json:"my_point"`
	TeamPoint    int    `json:"team_point"`
	MonthlyPoint int    `json:"monthly_point"`
	TeamMember   TeamMember
}

type TeamMember struct {
	HigherPearl   int `json:"higher_pearl"`
	HigherEmerald int `json:"higher_emerald"`
	HigherRuby    int `json:"higher_ruby"`
}

func VerifyLevel(memberId int) bool {
	member := FindMember(memberId)
	if checkCondition(member) {
		return promote(member)
	}
	return false
}

func checkCondition(member Member) bool {
	if member.Level == levelEmeraldAlpha &&
		member.MonthlyPoint >= conditionMonthlyPointOfRubyPup &&
		member.TeamPoint > conditionTeamPointOfRubyPup &&
		member.TeamMember.HigherEmerald >= conditionTeamMemberHigherEmeraldOfRubyPup {
		return true
	}
	return false
}

func countTeamMember(memberId int) TeamMember {
	db := mlmDB.Connect()
	defer db.Close()
	statement, err := db.Prepare(`SELECT COUNT(id) AS HigherThanLevel FROM members WHERE leader_id=? AND level >= ?`)
	if err != nil {
		panic(err.Error())
	}
	resultCountCountHigherPearlPup := statement.QueryRow(memberId, levelPearlPup)
	resultCountHigherEmeraldPup := statement.QueryRow(memberId, levelEmeraldPup)
	resultCountHigherRubyPup := statement.QueryRow(memberId, levelRubyPup)

	var teamMember TeamMember

	err = resultCountCountHigherPearlPup.Scan(&teamMember.HigherPearl)
	if err != nil {
		panic(err.Error())
	}

	err = resultCountHigherEmeraldPup.Scan(&teamMember.HigherEmerald)
	if err != nil {
		panic(err.Error())
	}

	err = resultCountHigherRubyPup.Scan(&teamMember.HigherRuby)
	if err != nil {
		panic(err.Error())
	}

	return teamMember
}

func FindMember(memberId int) Member {
	db := mlmDB.Connect()
	defer db.Close()
	statement, err := db.Prepare(`SELECT * FROM members WHERE id=?`)
	if err != nil {
		panic(err)
	}
	row := statement.QueryRow(memberId)
	var member Member
	err = row.Scan(
		&member.MemberId,
		&member.MemberName,
		&member.LeaderId,
		&member.Level,
	)
	if err != nil {
		panic(err.Error())
	}
	member.MyPoint = getMyPoint(memberId)
	member.MonthlyPoint = getMonthlyPoint(memberId, int(time.Now().Month()), time.Now().Year())
	member.TeamMember = countTeamMember(memberId)
	member.TeamPoint = getTeamPoint(memberId)
	return member
}

func promote(member Member) bool {
	db := mlmDB.Connect()
	defer db.Close()
	statement, _ := db.Prepare(`UPDATE members SET level=? WHERE id=?`)
	member.Level++
	_, err := statement.Exec(member.Level, member.MemberId)
	if err != nil {
		panic(err.Error())
	}
	return true
}

func demote(member Member) bool {
	db := mlmDB.Connect()
	defer db.Close()
	statement, _ := db.Prepare(`UPDATE members SET level=? WHERE id=?`)
	member.Level--
	_, err := statement.Exec(member.Level, member.MemberId)
	if err != nil {
		panic(err.Error())
	}
	return true
}

func getMonthlyPoint(memberId, month, year int) int {
	db := mlmDB.Connect()
	defer db.Close()
	statement, err := db.Prepare(`SELECT SUM(point) AS monthly_point FROM point_records WHERE member_id = ?  AND MONTH(create_date) = ? AND YEAR(create_date) = ?`)
	if err != nil {
		log.Println(err)
		return 0
	}
	row := statement.QueryRow(memberId, month, year)
	var monthlyPoint int
	err = row.Scan(&monthlyPoint)
	if err != nil {
		log.Println(err)
		return 0
	}
	return monthlyPoint
}

func getMyPoint(memberId int) int {
	db := mlmDB.Connect()
	defer db.Close()
	statement, err := db.Prepare(`SELECT SUM(point) AS my_point FROM point_records WHERE member_id = ?`)
	if err != nil {
		log.Println(err)
		return 0
	}
	row := statement.QueryRow(memberId)
	var myPoint int
	err = row.Scan(&myPoint)
	if err != nil {
		log.Println(err)
		return 0
	}
	return myPoint
}

func getTeamPoint(memberId int) int {
	db := mlmDB.Connect()
	defer db.Close()
	statement, err := db.Prepare(`SELECT SUM(point_records.point) AS team_point FROM point_records INNER JOIN members ON point_records.member_id = members.id WHERE members.leader_id = ? OR point_records.member_id = ?`)
	if err != nil {
		log.Println(err)
		return 0
	}
	row := statement.QueryRow(memberId, memberId)
	var teamPoint int
	err = row.Scan(&teamPoint)
	if err != nil {
		log.Println(err)
		return 0
	}
	return teamPoint
}
