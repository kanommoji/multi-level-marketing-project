package point

import (
	"database/sql"
	"multi-level-marketing-project/database/repository"
	"multi-level-marketing-project/internal/members"
)

func RecordPoint(db *sql.DB, newUserPoint members.NewUserPoint) bool {
	if !repository.Insert(db, newUserPoint) {
		return false
	}
	return true
}
