package event

import (
	"github.com/jmoiron/sqlx"
	"prc_hub_back/application/user"
	"prc_hub_back/domain/model/event"
)

type (
	CreateEventDocumentParam event.CreateEventDocumentParam
	UpdateEventDocumentParam event.UpdateEventDocumentParam
	GetDocumentQueryParam    event.GetDocumentQueryParam
)

func CreateDocument(db *sqlx.DB, p CreateEventDocumentParam, requestUserId int64) (_ event.EventDocument, err error) {
	// リクエスト元のユーザーを取得
	u, err := user.Get(requestUserId)
	if err != nil {
		return
	}

	return event.CreateEventDocument(
		db,
		event.CreateEventDocumentParam{
			EventId: p.EventId,
			Name:    p.Name,
			Url:     p.Url,
		},
		u,
	)
}

func GetDocument(db *sqlx.DB, id int64, requestUserId int64) (_ event.EventDocument, err error) {
	// リクエスト元のユーザーを取得
	u, err := user.Get(requestUserId)
	if err != nil {
		return
	}

	return event.GetDocument(
		db,
		id,
		u,
	)
}

func GetDocumentList(db *sqlx.DB, q GetDocumentQueryParam, requestUserId int64) ([]event.EventDocument, error) {
	return event.GetDocumentList(
		db,
		event.GetDocumentQueryParam{
			EventId:     q.EventId,
			Name:        q.Name,
			NameContain: q.NameContain,
		},
	)
}

func UpdateDocument(db *sqlx.DB, id int64, p UpdateEventDocumentParam, requestUserId int64) (event.EventDocument, error) {
	// リクエスト元のユーザーを取得
	u, err := user.Get(requestUserId)
	if err != nil {
		return event.EventDocument{}, err
	}

	return event.UpdateEventDocument(
		db,
		id,
		event.UpdateEventDocumentParam{
			Name: p.Name,
			Url:  p.Url,
		},
		u,
	)
}

func DeleteDocument(db *sqlx.DB, id int64, requestUserId int64) error {
	// リクエスト元のユーザーを取得
	u, err := user.Get(requestUserId)
	if err != nil {
		return err
	}

	return event.DeleteEventDocument(
		db,
		id,
		u,
	)
}
