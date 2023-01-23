package eisucon

import (
	"context"
	"github.com/jmoiron/sqlx"
)

func Migrate(db *sqlx.DB, sqlFile string) error {
	_, err := db.ExecContext(context.Background(), sqlFile)
	if err != nil {
		return err
	}
	return nil
}
