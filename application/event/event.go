package event

import (
	"github.com/jmoiron/sqlx"
	"prc_hub_back/application/user"
	"prc_hub_back/domain/model/event"
	userDomain "prc_hub_back/domain/model/user"
)

type (
	CreateEventParam       = event.CreateEventParam
	UpdateEventParam       = event.UpdateEventParam
	GetEventListQueryParam = event.GetEventListQueryParam
	GetEventQueryParam     = event.GetEventQueryParam
)

func CreateEvent(db *sqlx.DB, p CreateEventParam, requestUserId int64) (event.Event, error) {
	// リクエスト元のユーザーを取得
	u, err := user.Get(db, requestUserId)
	if err != nil {
		return event.Event{}, err
	}

	return event.CreateEvent(db, p, u)
}

func GetEvent(db *sqlx.DB, id int64, q GetEventQueryParam, requestUserId *int64) (event.EventEmbed, error) {
	u := new(userDomain.User)

	if requestUserId != nil {
		// リクエスト元のユーザーを取得
		var u2 userDomain.User
		u2, err := user.Get(db, *requestUserId)
		if err != nil {
			return event.EventEmbed{}, err
		}
		u = &u2
	} else if requestUserId == nil {
		// リクエストユーザーが指定されていない場合は最小権限のユーザーを仮使用
		u = &userDomain.User{
			Id:                  0,
			PostEventAvailabled: false,
			Manage:              false,
			Admin:               false,
		}
	}

	return event.GetEvent(db, id, q, *u)
}

func GetEventList(q GetEventListQueryParam, db *sqlx.DB, requestUserId *int64) ([]event.EventEmbed, error) {
	u := new(userDomain.User)

	if requestUserId != nil {
		// リクエスト元のユーザーを取得
		var u2 userDomain.User
		u2, err := user.Get(db, *requestUserId)
		if err != nil {
			return nil, err
		}
		u = &u2
	} else if requestUserId == nil {
		// リクエストユーザーが指定されていない場合は最小権限のユーザーを仮使用
		u = &userDomain.User{
			Id:                  0,
			PostEventAvailabled: false,
			Manage:              false,
			Admin:               false,
		}
	}

	return event.GetEventList(db, q, *u)
}

func UpdateEvent(db *sqlx.DB, id int64, p UpdateEventParam, requestUserId int64) (event.Event, error) {
	// リクエスト元のユーザーを取得
	u, err := user.Get(db, requestUserId)
	if err != nil {
		return event.Event{}, err
	}

	return event.UpdateEvent(db, id, p, u)
}

func DeleteEvent(db *sqlx.DB, id int64, requestUserId int64) error {
	// リクエスト元のユーザーを取得
	u, err := user.Get(db, requestUserId)
	if err != nil {
		return err
	}

	return event.DeleteEvent(db, id, u)
}
