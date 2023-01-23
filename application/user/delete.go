package user

import (
	"github.com/jmoiron/sqlx"
	"prc_hub_back/domain/model/user"
)

func Delete(db *sqlx.DB, id int64, requestUserId int64) error {
	// リクエスト元のユーザーを取得
	u, err := Get(db, id)
	if err != nil {
		return err
	}

	return user.DeleteUesr(db, id, u)
}
