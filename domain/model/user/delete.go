package user

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

// Errors
var (
	ErrAdminUserCannnotDelete = errors.New("cannot delete admin user")
)

func DeleteUesr(db *sqlx.DB, id int64, requestUser User) error {
	// リポジトリから削除対象の`User`を取得
	u, err := Get(db, id)
	if err != nil {
		return err
	}

	if !requestUser.Admin && requestUser.Id != id {
		// Admin権限なし 且つ IDが自分ではない場合は削除不可
		return ErrUserNotFound
	}

	if u.Admin {
		// Adminユーザーは削除不可
		return ErrAdminUserCannnotDelete
	}

	// `id`が一致する行を`users`テーブルから削除
	r2, err := db.Exec(
		`DELETE FROM users WHERE id = ?`,
		id,
	)
	if err != nil {
		return err
	}
	i, err := r2.RowsAffected()
	if err != nil {
		return err
	}
	if i != 1 {
		// 削除された行数が1ではない場合
		// `id`に一致する`uesr`が存在しない
		return ErrUserNotFound
	}
	return nil
}
