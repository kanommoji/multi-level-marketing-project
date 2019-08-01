package members

const (
	levelEmeraldAlpha = 6

	conditionMonthlyPointOfRubyPup = 400

	conditionTeamPointOfRubyPup = 20000

	conditionTeamMemberHigherEmerald = 2
)

func CheckCondition(member Member) bool {
	if member.Level == levelEmeraldAlpha &&
		member.MonthlyPoint >= conditionMonthlyPointOfRubyPup &&
		member.TeamPoint > conditionTeamPointOfRubyPup &&
		member.TeamMember.HigherEmerald >= 2 {
		return true
	}
	return false
}

type Member struct {
	MemberID     int
	MemberName   string
	LeaderID     int
	Level        int
	MyPoint      int
	MonthlyPoint int
	TeamPoint    int
	TeamMember   TeamMember
}

type TeamMember struct {
	HigherPearl   int
	HigherEmerald int
	HigherRuby    int
}

type NewUserPoint struct {
	UserRefferal int
	NewPoint     int
}
