package user

import (
	"context"
	"github.com/jmoiron/sqlx"
	"prc_hub_back/domain/model/sqlc"
)

func Get(db *sqlx.DB, id int64) (User, error) {
	queries := sqlc.New(db)

	u, _ := queries.GetUser(
		context.Background(),
		sqlc.GetUserParams{SetEmail: "%"},
	)

	// user
	user := User{
		Id:                  int64(u.ID),
		Name:                u.Name,
		Email:               u.Email,
		Password:            u.Password,
		PostEventAvailabled: u.PostEventAvailabled,
		Manage:              u.Manage,
		Admin:               u.Admin,
		TwitterId:           &u.TwitterID.String,
		GithubUsername:      &u.GithubUsername.String,
		StarCount:           uint64(u.StarCount),
	}

	return user, nil
}

func GetByEmail(db *sqlx.DB, email string) (User, error) {
	queries := sqlc.New(db)

	u, _ := queries.GetUserByVerify(
		context.Background(),
		email,
	)

	// user
	user := User{
		Id:       int64(u.ID),
		Email:    u.Email,
		Password: u.Password,
		Admin:    u.Admin,
	}

	return user, nil
}
