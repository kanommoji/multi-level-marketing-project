package member

import (
	"database/sql"
	"multi-level-marketing-project/config"
	"multi-level-marketing-project/model"
	repository "multi-level-marketing-project/repository/member"
)

const (
	levelPearlPup            = 1
	levelPearlJuvenile       = 2
	levelPearlAlpha          = 3
	levelEmeraldPup          = 4
	levelEmeraldJuvenile     = 5
	levelEmeraldAlpha        = 6
	levelRubyPup             = 7
	levelRubyJuvenile        = 8
	levelBlueDiamondJuvenile = 11

	conditionMyPointOfPearlJuvenile = 600
	conditionMyPointOfPearlAlpha    = 1000

	conditionMonthlyPointOfEmeraldPup       = 100
	conditionMonthlyPointOfEmeraldJuvenile  = 160
	conditionMonthlyPointOfEmeraldAlpha     = 200
	conditionMonthlyPointOfRubyPup          = 400
	conditionMonthlyPointOfRubyJuvenile     = 600
	conditionMonthlyPointOfRubyAlpha        = 800
	conditionMonthlyPointOfBlueDiamondAlpha = 2000

	conditionTeamPointOfEmeraldPup       = 4000
	conditionTeamPointOfEmeraldJuvenile  = 8000
	conditionTeamPointOfEmeraldAlpha     = 12000
	conditionTeamPointOfRubyPup          = 20000
	conditionTeamPointOfRubyJuvenile     = 40000
	conditionTeamPointOfRubyAlpha        = 100000
	conditionTeamPointOfBlueDiamondAlpha = 200000

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
	if member.Level == levelPearlPup &&
		member.MyPoint > conditionMyPointOfPearlJuvenile {
		return true
	}
	if member.Level == levelPearlJuvenile &&
		member.MyPoint > conditionMyPointOfPearlAlpha {
		return true
	}
	if member.Level == levelPearlAlpha &&
		member.MonthlyPoint >= conditionMonthlyPointOfEmeraldPup &&
		member.TeamPoint > conditionTeamPointOfEmeraldPup &&
		member.TeamMember.HigherPearl >= conditionTeamMemberHigherPearl {
		return true
	}
	if member.Level == levelEmeraldPup &&
		member.MonthlyPoint >= conditionMonthlyPointOfEmeraldJuvenile &&
		member.TeamPoint > conditionTeamPointOfEmeraldJuvenile {
		return true
	}
	if member.Level == levelEmeraldJuvenile &&
		member.MonthlyPoint >= conditionMonthlyPointOfEmeraldAlpha &&
		member.TeamPoint > conditionTeamPointOfEmeraldAlpha {
		return true
	}
	if member.Level == levelEmeraldAlpha &&
		member.MonthlyPoint >= conditionMonthlyPointOfRubyPup &&
		member.TeamPoint > conditionTeamPointOfRubyPup &&
		member.TeamMember.HigherEmerald >= conditionTeamMemberHigherEmerald {
		return true
	}
	if member.Level == levelRubyPup &&
		member.MonthlyPoint >= conditionMonthlyPointOfRubyJuvenile &&
		member.TeamPoint > conditionTeamPointOfRubyJuvenile {
		return true
	}
	if member.Level == levelRubyJuvenile &&
		member.MonthlyPoint >= conditionMonthlyPointOfRubyAlpha &&
		member.TeamPoint > conditionTeamPointOfRubyAlpha {
		return true
	}
	if member.Level == levelBlueDiamondJuvenile &&
		member.MonthlyPoint >= conditionMonthlyPointOfBlueDiamondAlpha &&
		member.TeamPoint > conditionTeamPointOfBlueDiamondAlpha {
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
	return repository.UpdateLevelPlusOne(database, memberID)
}
