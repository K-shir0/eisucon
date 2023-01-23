package event

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
	"prc_hub_back/domain/model/sqlc"
	"prc_hub_back/domain/model/user"
)

type GetEventQueryParam struct {
	Embed *[]string `query:"embed"`
}

func GetEvent(db *sqlx.DB, id int64, q GetEventQueryParam, requestUser user.User) (EventEmbed, error) {
	embedUser := false
	embedDocuments := false
	if q.Embed != nil {
		for _, e := range *q.Embed {
			if e == "user" {
				embedUser = true
			}
			if e == "documents" {
				embedDocuments = true
			}
		}
	}

	if embedUser && embedDocuments {
		event, err := ConvEventByWithUserAndDocuments(sqlc.New(db), id)
		if err != nil {
			return EventEmbed{}, err
		}

		return *event, nil
	}

	if embedUser {
		event, err := ConvEventByWithUser(sqlc.New(db), id)
		if err != nil {
			return EventEmbed{}, err
		}

		return *event, nil
	}

	if embedDocuments {
		event, err := ConvEventByWithDocuments(sqlc.New(db), id)
		if err != nil {
			return EventEmbed{}, err
		}

		return *event, nil
	}

	event, err := ConvEvent(sqlc.New(db), id)
	if err != nil {
		return EventEmbed{}, err
	}

	return *event, nil
}

// ConvEventByWithUserAndDocuments は、Event と User と Document を取得する
func ConvEventByWithUserAndDocuments(queries *sqlc.Queries, event_id int64) (*EventEmbed, error) {
	eventsRows, err := queries.GetEventWithUserAndDocuments(context.Background(), int32(event_id))
	if err != nil {
		return nil, err
	}
	if len(eventsRows) == 0 {
		return nil, nil
	}

	var event *EventEmbed
	var documents []EventDocument

	for _, row := range eventsRows {
		if event == nil {
			event = &EventEmbed{
				Event: Event{
					Id:          int64(row.ID),
					Name:        row.Name,
					Description: &row.Description.String,
					Location:    &row.Location.String,
					Datetimes:   []EventDatetime{},
					Published:   row.Published,
					Completed:   row.Completed,
					UserId:      int64(row.UserID),
				},
				User: &user.User{
					Id:                  int64(row.UserID),
					Name:                row.Name_2,
					Email:               row.Email,
					Password:            row.Password,
					PostEventAvailabled: row.PostEventAvailabled,
					Manage:              row.Manage,
					Admin:               row.Admin,
					TwitterId:           &row.TwitterID.String,
					GithubUsername:      &row.GithubUsername.String,
					StarCount:           uint64(row.StarCount),
				},
			}
		}

		// Time
		eventDatetime := &EventDatetime{
			Start: row.Start,
			End:   row.End,
		}

		// Document
		eventDocuments := &EventDocument{
			EventId: int64(row.DocumentEventID),
			Id:      int64(row.DocumentID),
			Name:    row.DocumentName,
			Url:     row.Url,
		}

		// Time の処理
		event.Event.Datetimes = append(event.Event.Datetimes, *eventDatetime)
		documents = append(documents, *eventDocuments)
	}

	// 上と重複
	// Time の処理
	if event != nil {
		event.Event.Datetimes = lo.Uniq[EventDatetime](event.Event.Datetimes)
		// Event の処理
		uniqDocuments := lo.Uniq[EventDocument](documents)
		event.Documents = &uniqDocuments
	}

	return event, nil
}

// ConvEventByWithUser は、Event と User を取得する
func ConvEventByWithUser(queries *sqlc.Queries, event_id int64) (*EventEmbed, error) {
	eventsRows, err := queries.GetEventWithDocuments(context.Background(), int32(event_id))
	if err != nil {
		return nil, err
	}
	if len(eventsRows) == 0 {
		return nil, nil
	}

	var event *EventEmbed

	for _, row := range eventsRows {
		if event == nil {
			event = &EventEmbed{
				Event: Event{
					Id:          int64(row.ID),
					Name:        row.Name,
					Description: &row.Description.String,
					Location:    &row.Location.String,
					Datetimes:   []EventDatetime{},
					Published:   row.Published,
					Completed:   row.Completed,
					UserId:      int64(row.UserID),
				},
			}
		}

		// Time
		eventDatetime := &EventDatetime{
			Start: row.Start,
			End:   row.End,
		}

		// Time の処理
		event.Event.Datetimes = append(event.Event.Datetimes, *eventDatetime)
	}

	// 上と重複
	// Time の処理
	if event != nil {
		event.Event.Datetimes = lo.Uniq[EventDatetime](event.Event.Datetimes)
	}

	return event, nil
}

// ConvEventByWithDocuments は、Event と Document を取得する
func ConvEventByWithDocuments(queries *sqlc.Queries, event_id int64) (*EventEmbed, error) {
	eventsRows, err := queries.GetEventWithDocuments(context.Background(), int32(event_id))
	if err != nil {
		return nil, err
	}
	if len(eventsRows) == 0 {
		return nil, nil
	}

	var event *EventEmbed
	var documents []EventDocument

	for _, row := range eventsRows {
		if event == nil {
			event = &EventEmbed{
				Event: Event{
					Id:          int64(row.ID),
					Name:        row.Name,
					Description: &row.Description.String,
					Location:    &row.Location.String,
					Datetimes:   []EventDatetime{},
					Published:   row.Published,
					Completed:   row.Completed,
					UserId:      int64(row.UserID),
				},
			}
		}

		// Time
		eventDatetime := &EventDatetime{
			Start: row.Start,
			End:   row.End,
		}

		// Document
		eventDocuments := &EventDocument{
			EventId: int64(row.DocumentEventID),
			Id:      int64(row.DocumentID),
			Name:    row.DocumentName,
			Url:     row.Url,
		}

		// Time の処理
		event.Event.Datetimes = append(event.Event.Datetimes, *eventDatetime)
		documents = append(documents, *eventDocuments)
	}

	// 上と重複
	// Time の処理
	if event != nil {
		event.Event.Datetimes = lo.Uniq[EventDatetime](event.Event.Datetimes)
		// Event の処理
		uniqDocuments := lo.Uniq[EventDocument](documents)
		event.Documents = &uniqDocuments
	}

	return event, nil
}

// ConvEvent は、Event と User と Document を取得する
func ConvEvent(queries *sqlc.Queries, event_id int64) (*EventEmbed, error) {
	eventsRows, err := queries.GetEvent(context.Background(), int32(event_id))
	if err != nil {
		return nil, err
	}
	if len(eventsRows) == 0 {
		return nil, nil
	}

	var event *EventEmbed

	for _, row := range eventsRows {
		if event == nil {
			event = &EventEmbed{
				Event: Event{
					Id:          int64(row.ID),
					Name:        row.Name,
					Description: &row.Description.String,
					Location:    &row.Location.String,
					Datetimes:   []EventDatetime{},
					Published:   row.Published,
					Completed:   row.Completed,
					UserId:      int64(row.UserID),
				},
			}
		}

		// Time
		eventDatetime := &EventDatetime{
			Start: row.Start,
			End:   row.End,
		}

		// Time の処理
		event.Event.Datetimes = append(event.Event.Datetimes, *eventDatetime)
	}

	// 上と重複
	// Time の処理
	if event != nil {
		event.Event.Datetimes = lo.Uniq[EventDatetime](event.Event.Datetimes)
	}

	return event, nil
}
