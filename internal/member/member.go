package member

import (
	"database/sql"
	"multi-level-marketing-project/config"
	"multi-level-marketing-project/model"
	repository "multi-level-marketing-project/repository/member"
)

func FindMember(database *sql.DB, memberID int) model.Member {
	member := repository.FindMemberByID(database, memberID)
	member.MyPoint = repository.GetMyPoint(database, memberID)
	member.MonthlyPoint = repository.GetMonthlyPoint(database, memberID, int(config.TimeNow().Month()), config.TimeNow().Year())
	member.TeamPoint = repository.GetTeamPoint(database, memberID)
	member.TeamMember = repository.CountTeamMember(database, memberID)
	return member
}
