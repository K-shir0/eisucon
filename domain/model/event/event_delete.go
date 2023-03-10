package event

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"prc_hub_back/domain/model/user"
)

// Errors
var (
	ErrCannotDeleteEvent = errors.New("sorry, you cannot delete this event")
)

func DeleteEvent(db *sqlx.DB, id int64, requestUser user.User) error {
	// Get event
	e, err := GetEvent(
		db,
		id,
		GetEventQueryParam{},
		requestUser,
	)
	if err != nil {
		return err
	}

	// 権限の検証
	if !requestUser.Admin && !requestUser.Manage && requestUser.Id != e.UserId {
		// Admin権限なし 且つ `Event.UserId`が自分ではない場合は削除不可
		return ErrCannotDeleteEvent
	}

	// `id`が一致する行を`events`テーブルから削除
	r2, err := db.Exec(
		`DELETE FROM events WHERE id = ?`,
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
		// 変更された行数が1ではない場合
		// `id`に一致する`document`が存在しない
		return ErrEventNotFound
	}

	return nil
}
