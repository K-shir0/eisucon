package user

import (
	"github.com/jmoiron/sqlx"
	"prc_hub_back/domain/model/user"
)

type (
	UpdateUserParam user.UpdateUserParam
)

func Update(db *sqlx.DB, id int64, p UpdateUserParam, requestUserId int64) (user.UserWithToken, error) {
	// リクエスト元のユーザーを取得
	u, err := Get(db, requestUserId)
	if err != nil {
		return user.UserWithToken{}, err
	}

	return user.Update(
		db,
		id,
		user.UpdateUserParam{
			Name:                p.Name,
			Email:               p.Email,
			Password:            p.Password,
			PostEventAvailabled: p.PostEventAvailabled,
			Manage:              p.Manage,
			TwitterId:           p.TwitterId,
			GithubUsername:      p.GithubUsername,
		},
		u,
	)
}
