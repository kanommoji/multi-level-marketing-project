package members

const (
	levelPearlPup     = 1
	levelPearlAlpha   = 3
	levelEmeraldAlpha = 6

	conditionMyPointOfPearlJuvenile = 600

	conditionMonthlyPointOfEmeraldPup = 100
	conditionMonthlyPointOfRubyPup    = 400

	conditionTeamPointOfEmeraldPup = 4000
	conditionTeamPointOfRubyPup    = 20000

	conditionTeamMemberHigherPearl   = 2
	conditionTeamMemberHigherEmerald = 2
)

func CheckCondition(member Member) bool {
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
