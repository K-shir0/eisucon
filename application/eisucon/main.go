package eisucon

import (
	"github.com/jmoiron/sqlx"
	"prc_hub_back/domain/model/eisucon"
)

// Singleton field
var migrateSqlFile string

func Migrate(db *sqlx.DB) error {
	return eisucon.Migrate(db, migrateSqlFile)
}
