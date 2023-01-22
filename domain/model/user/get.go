package user

import (
	"context"
	"prc_hub_back/domain/model/sqlc"
)

func Get(id int64) (User, error) {
	// MySQLサーバーに接続
	db, err := OpenMysql()
	if err != nil {
		return User{}, err
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer db.Close()

	queries := sqlc.New(db)

	u, err := queries.GetUser(
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

func GetByEmail(email string) (User, error) {
	// MySQLサーバーに接続
	db, err := OpenMysql()
	if err != nil {
		return User{}, err
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer db.Close()

	queries := sqlc.New(db)

	u, err := queries.GetUser(
		context.Background(),
		sqlc.GetUserParams{SetEmail: email},
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
