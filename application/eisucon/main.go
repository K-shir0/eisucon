package eisucon

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"prc_hub_back/domain/model/eisucon"
)

// Singleton field
var db *sqlx.DB
var migrate string

func Init(db2 *sqlx.DB, migratedata string) {
	db = db2
	migrate = migratedata
}

func Migrate() error {
	if migrate == "" {
		return errors.New("migrate sql file does not set")
	}
	return eisucon.Migrate(db, migrate)
}
