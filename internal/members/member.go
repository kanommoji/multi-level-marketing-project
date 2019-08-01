package members

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
	UserRefferal int `json:"user_referral"`
	NewPoint     int `json:"new_point"`
}

func VerifyLevel(MemberID int) bool {
	return true
}
