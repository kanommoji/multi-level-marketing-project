package member

import (
	"database/sql"
	"multi-level-marketing-project/config"
	"multi-level-marketing-project/model"
	repository "multi-level-marketing-project/repository/member"
)

const (
	levelPearlPup    = 1
	levelPearlAlpha  = 3
	levelEmeralAlpha = 6

	conditionMyPointOfPearlJuvenile = 600

	conditionMonthlyPointOfEmeraldPup = 100
	conditionMonthlyPointOfRubyPup    = 400

	conditionTeamPointOfEmeraldPup = 4000
	conditionTeamPointOfRubyPup    = 20000

	conditionTeamMemberHigherPearl   = 2
	conditionTeamMemberHigherEmerald = 2
)

func FindMember(database *sql.DB, memberID int) model.Member {
	member := repository.FindMemberByID(database, memberID)
	member.MyPoint = repository.GetMyPoint(database, memberID)
	member.MonthlyPoint = repository.GetMonthlyPoint(database, memberID, int(config.TimeNow().Month()), config.TimeNow().Year())
	member.TeamPoint = repository.GetTeamPoint(database, memberID)
	member.TeamMember = repository.CountTeamMember(database, memberID)
	return member
}

func CheckCondition(member model.Member) bool {
	if member.Level == levelPearlPup && member.MyPoint > conditionMyPointOfPearlJuvenile {
		return true
	}
	if member.Level == levelPearlAlpha &&
		member.MonthlyPoint >= conditionMonthlyPointOfEmeraldPup &&
		member.TeamPoint > conditionTeamPointOfEmeraldPup &&
		member.TeamMember.HigherPearl >= conditionTeamMemberHigherPearl {
		return true
	}
	if member.Level == levelEmeralAlpha &&
		member.MonthlyPoint >= conditionMonthlyPointOfRubyPup &&
		member.TeamPoint > conditionTeamPointOfRubyPup &&
		member.TeamMember.HigherEmerald >= conditionTeamMemberHigherEmerald {
		return true
	}
	return false
}

func VerifyLevel(database *sql.DB, memberID int) bool {
	member := FindMember(database, memberID)
	if CheckCondition(member) {
		return Promote(database, memberID)
	}
	return false
}

func Promote(database *sql.DB, memberID int) bool {
	if !repository.UpdateLevelPlusOne(database, memberID) {
		return false
	}
	return true
}
