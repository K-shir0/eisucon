package user

import (
	"github.com/jmoiron/sqlx"
	"prc_hub_back/domain/model/user"
)

type GetUserListQuery user.GetUserListQueryParam

func GetList(db *sqlx.DB, q GetUserListQuery) ([]user.User, error) {
	return user.GetList(db, user.GetUserListQueryParam(q))
}
