package models

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
