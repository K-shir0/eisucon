package user

import (
	"github.com/jmoiron/sqlx"
	"prc_hub_back/domain/model/user"
)

func AddStar(db *sqlx.DB, userId uint64) (count uint64, err error) {
	return user.AddStar(db, userId)
}
