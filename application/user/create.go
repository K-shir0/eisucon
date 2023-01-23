package user

import (
	"github.com/jmoiron/sqlx"
	"prc_hub_back/domain/model/user"
)

type (
	CreateUserParam user.CreateUserParam
)

func Create(db *sqlx.DB, p CreateUserParam) (user.UserWithToken, error) {
	return user.CreateUser(
		db,
		user.CreateUserParam{
			Name:           p.Name,
			Email:          p.Email,
			Password:       p.Password,
			TwitterId:      p.TwitterId,
			GithubUsername: p.GithubUsername,
		},
	)
}
