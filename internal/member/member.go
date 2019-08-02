package member

import "multi-level-marketing-project/model"

func FindMember(memberID int) model.Member {
	return model.Member{
		MemberID:     10029,
		MemberName:   "ชนา",
		LeaderID:     20029,
		Level:        6,
		MyPoint:      1000,
		MonthlyPoint: 350,
		TeamPoint:    20000,
	}
}
