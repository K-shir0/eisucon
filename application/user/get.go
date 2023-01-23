package user

import (
	"github.com/jmoiron/sqlx"
	"prc_hub_back/domain/model/user"
)

func Get(db *sqlx.DB, id int64) (_ user.User, err error) {
	return user.Get(db, id)
}

func GetByEmail(db *sqlx.DB, email string) (user.User, error) {
	return user.GetByEmail(db, email)
}
