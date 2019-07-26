package members

import (
	mlmDB "multi-level-marketing-project/database/db"
	"log"
	"time"
)

const (
	PearlPup                                       = 1
	EmeraldPup                                     = 4
	EmeraldAlpha                                   = 6
	RubyPup                                        = 7
	conditionMonthlyPointOfEmeraldAlpha            = 400
	conditionTeamPointOfEmeraldAlpha               = 20000
	conditionTeamMemberHigherEmeraldOfEmeraldAlpha = 2
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

func checkCondition(member Member) bool {
	if member.Level == EmeraldAlpha && member.MonthlyPoint >= conditionMonthlyPointOfEmeraldAlpha && member.TeamPoint > conditionTeamPointOfEmeraldAlpha && member.TeamMember.HigherEmerald >= conditionTeamMemberHigherEmeraldOfEmeraldAlpha {
		return true
	}
	return false
}

func CountTeamMember(memberId int) TeamMember {
	db := mlmDB.Connect()
	defer db.Close()
	statement, _ := db.Prepare("SELECT COUNT(id) AS HigherPearl FROM mlm.members WHERE leader_id=? AND level >= ?")
	resultCountCountHigherPearlPup, _ := statement.Query(memberId, PearlPup)
	resultCountHigherEmeraldPup, _ := statement.Query(memberId, EmeraldPup)
	resultCountHigherRubyPup, _ := statement.Query(memberId, RubyPup)
	var teamMember TeamMember
	if resultCountCountHigherPearlPup.Next() {
		err := resultCountCountHigherPearlPup.Scan(&teamMember.HigherPearl)
		if err != nil {
			panic(err.Error())
		}
	}
	if resultCountHigherEmeraldPup.Next() {
		err := resultCountHigherEmeraldPup.Scan(&teamMember.HigherEmerald)
		if err != nil {
			panic(err.Error())
		}
	}
	if resultCountHigherRubyPup.Next() {
		err := resultCountHigherRubyPup.Scan(&teamMember.HigherRuby)
		if err != nil {
			panic(err.Error())
		}
	}
	return teamMember
}

func getMonthlyPoint(memberId, month, year int) int {
	database := mlmDB.Connect()
	statement, err := database.Prepare(`SELECT SUM(POINT) AS monthly_point FROM point_records WHERE member_id = ?  AND MONTH(create_date) = ? AND YEAR(create_date) = ?`)
	if err != nil {
		log.Println(err)
		return 0
	}
	row := statement.QueryRow(memberId, month, year)
	var monthlyPoint int
	row.Scan(&monthlyPoint)
	return monthlyPoint
}

func getCurrentMonth() int {
	return int(time.Now().Month())
}

func Promote(memberAlpha Member) bool {
	db := mlmDB.Connect()
	defer db.Close()
	statement, _ := db.Prepare("UPDATE mlm.members SET level=? WHERE id=?")
	memberAlpha.Level++
	_, err := statement.Exec(memberAlpha.Level, memberAlpha.MemberId)
	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}

func Demote(memberAlpha Member) bool {
	db := mlmDB.Connect()
	defer db.Close()
	statement, _ := db.Prepare("UPDATE mlm.members SET level=? WHERE id=?")
	memberAlpha.Level--
	_, err := statement.Exec(memberAlpha.Level, memberAlpha.MemberId)
	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}

func getMyPoint(memberId int) int {
	db2 := mlmDB.Connect()
	defer db2.Close()
	statement, err := db2.Prepare(`SELECT SUM(POINT) AS my_point FROM point_records WHERE member_id = ?`)
	if err != nil {
		log.Println(err)
		return 0
	}
	row := statement.QueryRow(memberId)
	var myPoint int
	row.Scan(&myPoint)
	return myPoint
}

func FindMember(memberId int) Member {
	db2 := mlmDB.Connect()
	defer db2.Close()
	statement, _ := db2.Prepare("SELECT * FROM mlm.members WHERE id=?")
	row, _ := statement.Query(memberId)
	var memberAlpha Member
	if row.Next() {
		err := row.Scan(
			&memberAlpha.MemberId,
			&memberAlpha.MemberName,
			&memberAlpha.LeaderId,
			&memberAlpha.Level,
		)
		if err != nil {
			panic(err.Error())
		}
	}
	memberAlpha.MyPoint = getMyPoint(memberId)
	memberAlpha.MonthlyPoint = getMonthlyPoint(memberId, getCurrentMonth(), time.Now().Year())
	memberAlpha.TeamMember = CountTeamMember(memberId)
	memberAlpha.TeamPoint = getTeamPoint(memberId)
	return memberAlpha
}

func getTeamPoint(memberId int) int {
	db2 := mlmDB.Connect()
	defer db2.Close()
	statement, err := db2.Prepare(`SELECT SUM(point_records.point) AS team_point FROM mlm.point_records INNER JOIN mlm.members ON point_records.member_id = members.id WHERE members.leader_id = ? OR point_records.member_id = ?`)
	if err != nil {
		log.Println(err)
		return 0
	}
	row := statement.QueryRow(memberId, memberId)
	var teamPoint int
	row.Scan(&teamPoint)
	return teamPoint
}

func VerifyLevel(memberId int) bool {
	memberAlpha := FindMember(memberId)
	if checkCondition(memberAlpha) {
		return Promote(memberAlpha)
	}
	return false
}
