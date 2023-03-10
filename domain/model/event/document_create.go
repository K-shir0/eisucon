package event

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"prc_hub_back/domain/model/user"
)

// Errors
var (
	ErrCannotCreateEventDocument = errors.New("sorry, you cannot create document to this event")
)

type CreateEventDocumentParam struct {
	EventId int64  `json:"event_id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
}

func (p CreateEventDocumentParam) validate(db *sqlx.DB, requestUser user.User) error {
	// フィールドの検証
	err := validateDocumentName(p.Name)
	if err != nil {
		return err
	}
	err = validateUrl(p.Url)
	if err != nil {
		return err
	}
	err = validateEventId(db, p.EventId, requestUser)
	if err != nil {
		return err
	}

	// 権限の検証
	if !requestUser.Admin && !requestUser.Manage {
		// Eventを取得
		e, err := GetEvent(db, p.EventId, GetEventQueryParam{}, requestUser)
		if err != nil {
			return err
		}

		if requestUser.Id != e.UserId {
			// `Admin`・`Manage`のいずれでもなく`Event.UserId`が自分ではない場合は追加不可
			return ErrCannotCreateEventDocument
		}
	}

	return nil
}

func CreateEventDocument(db *sqlx.DB, p CreateEventDocumentParam, requestUser user.User) (EventDocument, error) {
	err := p.validate(db, requestUser)
	if err != nil {
		return EventDocument{}, err
	}

	// `documents`テーブルに追加
	r, err := db.Exec(
		`INSERT INTO documents (event_id, name, url) VALUES (?, ?, ?)`,
		p.EventId, p.Name, p.Url,
	)
	if err != nil {
		return EventDocument{}, err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return EventDocument{}, err
	}
	e := EventDocument{
		Id:      id,
		EventId: p.EventId,
		Name:    p.Name,
		Url:     p.Url,
	}

	return e, nil
}
