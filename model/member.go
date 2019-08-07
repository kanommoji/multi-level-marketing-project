package model

type Member struct {
	MemberID     int        `json:"member_id"`
	MemberName   string     `json:"member_name"`
	LeaderID     int        `json:"leader_id"`
	Level        int        `json:"level"`
	MyPoint      int        `json:"my_point"`
	MonthlyPoint int        `json:"monthly_point"`
	TeamPoint    int        `json:"team_point"`
	TeamMember   TeamMember `json:"team_member"`
}

type TeamMember struct {
	HigherPearl   int `json:"higher_pearl"`
	HigherEmerald int `json:"higher_emerald"`
	HigherRuby    int `json:"higher_ruby"`
}
