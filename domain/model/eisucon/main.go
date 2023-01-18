package eisucon

import (
	"github.com/jmoiron/sqlx"
)

func Migrate(db *sqlx.DB, sqlFile string) error {
	_, err := sqlx.LoadFile(db, sqlFile)
	if err != nil {
		return err
	}
	return nil
}
