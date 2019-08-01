package point

import (
	"database/sql"
	"multi-level-marketing-project/database/repository"
	"multi-level-marketing-project/models"
)

func RecordPoint(db *sql.DB, newUserPoint models.NewUserPoint) bool {
	if !repository.Insert(db, newUserPoint) {
		return false
	}
	return true
}
