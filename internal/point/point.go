package point

import (
	"database/sql"
	"multi-level-marketing-project/model"
	repository "multi-level-marketing-project/repository/point"
)

func RecordPoint(database *sql.DB, newUserPoint model.NewUserPoint) bool {
	if !repository.Insert(database, newUserPoint) {
		return false
	}
	return true
}
