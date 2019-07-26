package members

import (
	"log"
	mlmDB "multi-level-marketing-project/database/db"
	"time"
)

const (
	PearlPup     = 1
	EmeraldPup   = 4
	EmeraldAlpha = 6
	RubyPup      = 7

	conditionMyPointOfPearlJuvenile           = 600
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

func checkCondition(member Member) bool {
	if member.Level == PearlPup && member.MyPoint > conditionMyPointOfPearlJuvenile {
		return true
	}
	if member.Level == EmeraldAlpha && member.MonthlyPoint >= conditionMonthlyPointOfRubyPup && member.TeamPoint > conditionTeamPointOfRubyPup && member.TeamMember.HigherEmerald >= conditionTeamMemberHigherEmeraldOfRubyPup {
		return true
	}
	return false
}

func CountTeamMember(memberId int) TeamMember {
	db := mlmDB.Connect()
	defer db.Close()
	statement, err := db.Prepare("SELECT COUNT(id) AS HigherThanLevel FROM mlm.members WHERE leader_id=? AND level >= ?")
	if err != nil {
		panic(err.Error())
	}
	resultCountCountHigherPearlPup := statement.QueryRow(memberId, PearlPup)
	resultCountHigherEmeraldPup := statement.QueryRow(memberId, EmeraldPup)
	resultCountHigherRubyPup := statement.QueryRow(memberId, RubyPup)
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

func Promote(member Member) bool {
	db := mlmDB.Connect()
	defer db.Close()
	statement, _ := db.Prepare("UPDATE mlm.members SET level=? WHERE id=?")
	member.Level++
	_, err := statement.Exec(member.Level, member.MemberId)
	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}

func Demote(member Member) bool {
	db := mlmDB.Connect()
	defer db.Close()
	statement, _ := db.Prepare("UPDATE mlm.members SET level=? WHERE id=?")
	member.Level--
	_, err := statement.Exec(member.Level, member.MemberId)
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
	var member Member
	if row.Next() {
		err := row.Scan(
			&member.MemberId,
			&member.MemberName,
			&member.LeaderId,
			&member.Level,
		)
		if err != nil {
			panic(err.Error())
		}
	}
	member.MyPoint = getMyPoint(memberId)
	member.MonthlyPoint = getMonthlyPoint(memberId, getCurrentMonth(), time.Now().Year())
	member.TeamMember = CountTeamMember(memberId)
	member.TeamPoint = getTeamPoint(memberId)
	return member
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
	member := FindMember(memberId)
	if checkCondition(member) {
		return Promote(member)
	}
	return false
}
