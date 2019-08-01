package members

import (
	"database/sql"
	"multi-level-marketing-project/database/repository"
	"multi-level-marketing-project/models"
	"time"
)

const (
	levelPearlPup            = 1
	levelPearlAlpha          = 3
	levelEmeraldAlpha        = 6
	levelBlueDiamondJuvenile = 11

	conditionMyPointOfPearlJuvenile = 600

	conditionMonthlyPointOfEmeraldPup       = 100
	conditionMonthlyPointOfRubyPup          = 400
	conditionMonthlyPointOfBlueDiamondAlpha = 2000

	conditionTeamPointOfEmeraldPup       = 4000
	conditionTeamPointOfRubyPup          = 20000
	conditionTeamPointOfBlueDiamondAlpha = 200000

	conditionTeamMemberHigherPearl   = 2
	conditionTeamMemberHigherEmerald = 2
)

func CheckCondition(member models.Member) bool {
	if member.Level == levelPearlPup &&
		member.MyPoint > conditionMyPointOfPearlJuvenile {
		return true
	}
	if member.Level == levelPearlAlpha &&
		member.MonthlyPoint >= conditionMonthlyPointOfEmeraldPup &&
		member.TeamPoint > conditionTeamPointOfEmeraldPup &&
		member.TeamMember.HigherPearl >= conditionTeamMemberHigherPearl {
		return true
	}
	if member.Level == levelEmeraldAlpha &&
		member.MonthlyPoint >= conditionMonthlyPointOfRubyPup &&
		member.TeamPoint > conditionTeamPointOfRubyPup &&
		member.TeamMember.HigherEmerald >= conditionTeamMemberHigherEmerald {
		return true
	}
	if member.Level == levelBlueDiamondJuvenile &&
		member.MonthlyPoint >= conditionMonthlyPointOfBlueDiamondAlpha &&
		member.TeamPoint > conditionTeamPointOfBlueDiamondAlpha {
		return true
	}
	return false
}

func FindMember(db *sql.DB, memberID int) models.Member {
	var member models.Member
	member = repository.GetMemberData(db, memberID)
	member.MyPoint = repository.GetMyPoint(db, memberID)
	member.MonthlyPoint = repository.GetMonthlyPoint(db, memberID,int(time.Now().Month()),time.Now().Year())
	member.TeamPoint = repository.GetTeamPoint(db, memberID)
	member.TeamMember = repository.CountTeamMember(db, memberID)
	return member
}
