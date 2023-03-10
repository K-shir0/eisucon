package event

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
	"prc_hub_back/domain/model/sqlc"
	"prc_hub_back/domain/model/user"
)

type GetEventListQueryParam struct {
	Published       *bool     `query:"published"`
	Name            *string   `query:"name"`
	NameContain     *string   `query:"name_contain"`
	Location        *string   `query:"location"`
	LocationContain *string   `query:"location_contain"`
	Embed           *[]string `query:"embed"`
}

func GetEventList(db *sqlx.DB, q GetEventListQueryParam, requestUser user.User) ([]EventEmbed, error) {
	queries := sqlc.New(db)

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

	var eventName = "%"
	var location = "%"
	var setNotPublished = true
	var published bool

	// クエリを作成
	if q.Name != nil {
		// イベント名の一致で絞り込み
		eventName = *q.Name
	}
	if q.NameContain != nil {
		// イベント名に文字列が含まれるかで絞り込み
		eventName = "%" + *q.NameContain + "%"
	}
	if q.Location != nil {
		// `Location`の一致で絞り込み
		location = *q.Location
	}
	if q.LocationContain != nil {
		// `Location`に文字列が含まれるかで絞り込み
		location = "%" + *q.LocationContain + "%"
	}
	if q.Published != nil {
		// `Published`で絞り込み
		setNotPublished = false
		published = *q.Published
	}

	if embedUser && embedDocuments {
		return ConvEventListByWithUserAndDocuments(queries, eventName, location, setNotPublished, published)
	}

	if embedUser {
		return ConvEventListByWithUser(queries, eventName, location, setNotPublished, published)
	}

	if embedDocuments {
		return ConvEventListByWithDocuments(queries, eventName, location, setNotPublished, published)
	}

	return ConvEventList(queries, eventName, location, setNotPublished, published)
}

// ConvEventListByWithUserAndDocuments は、Event と User と Document を取得する
func ConvEventListByWithUserAndDocuments(queries *sqlc.Queries, eventName string, location string, setNotPublished bool, published bool) ([]EventEmbed, error) {
	eventsRows, err := queries.ListEventsWithUserAndDocuments(context.Background(), sqlc.ListEventsWithUserAndDocumentsParams{
		SetEventName: eventName,
		SetLocation: sql.NullString{
			String: location,
			Valid:  true,
		},
		NotSetPublished: setNotPublished,
		SetPublished:    published,
	})
	if err != nil {
		return nil, err
	}
	if len(eventsRows) == 0 {
		return nil, nil
	}

	var prevEventId = eventsRows[0].ID
	var events []EventEmbed
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

		if prevEventId != row.ID {
			// Time の処理
			event.Event.Datetimes = lo.Uniq[EventDatetime](event.Event.Datetimes)
			// Event の処理
			uniqDocuments := lo.Uniq[EventDocument](documents)
			event.Documents = &uniqDocuments

			events = append(events, *event)
			event = nil
			documents = []EventDocument{}
		} else {
			// Time の処理
			event.Event.Datetimes = append(event.Event.Datetimes, *eventDatetime)
			documents = append(documents, *eventDocuments)
		}

		prevEventId = row.ID
	}

	// 上と重複
	// Time の処理
	if event != nil {
		event.Event.Datetimes = lo.Uniq[EventDatetime](event.Event.Datetimes)
		// Event の処理
		uniqDocuments := lo.Uniq[EventDocument](documents)
		event.Documents = &uniqDocuments

		events = append(events, *event)
	}

	return events, nil
}

// ConvEventListByWithUser User を含む
func ConvEventListByWithUser(queries *sqlc.Queries, eventName string, location string, setNotPublished bool, published bool) ([]EventEmbed, error) {
	eventsRows, err := queries.ListEventsWithUser(context.Background(), sqlc.ListEventsWithUserParams{
		SetEventName: eventName,
		SetLocation: sql.NullString{
			String: location,
			Valid:  true,
		},
		NotSetPublished: setNotPublished,
		SetPublished:    published,
	})
	if err != nil {
		return nil, err
	}
	if len(eventsRows) == 0 {
		return nil, nil
	}

	var prevEventId = eventsRows[0].ID
	var events []EventEmbed
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

		if prevEventId != row.ID {
			// Time の処理
			event.Event.Datetimes = lo.Uniq[EventDatetime](event.Event.Datetimes)

			events = append(events, *event)
			event = nil
		} else {
			// Time の処理
			event.Event.Datetimes = append(event.Event.Datetimes, *eventDatetime)
		}

		prevEventId = row.ID
	}

	// 上と重複
	// Time の処理
	if event != nil {
		event.Event.Datetimes = lo.Uniq[EventDatetime](event.Event.Datetimes)

		events = append(events, *event)
	}

	return events, nil
}

// ConvEventListByWithDocuments Documents を含む
func ConvEventListByWithDocuments(queries *sqlc.Queries, eventName string, location string, setNotPublished bool, published bool) ([]EventEmbed, error) {
	eventsRows, err := queries.ListEventsWithDocuments(context.Background(), sqlc.ListEventsWithDocumentsParams{
		SetEventName: eventName,
		SetLocation: sql.NullString{
			String: location,
			Valid:  true,
		},
		NotSetPublished: setNotPublished,
		SetPublished:    published,
	})
	if err != nil {
		return nil, err
	}
	if len(eventsRows) == 0 {
		return nil, nil
	}

	var prevEventId = eventsRows[0].ID
	var events []EventEmbed
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

		if prevEventId != row.ID {
			// Time の処理
			event.Event.Datetimes = lo.Uniq[EventDatetime](event.Event.Datetimes)
			// Event の処理
			uniqDocuments := lo.Uniq[EventDocument](documents)
			event.Documents = &uniqDocuments

			events = append(events, *event)
			event = nil
			documents = []EventDocument{}
		} else {
			// Time の処理
			event.Event.Datetimes = append(event.Event.Datetimes, *eventDatetime)
			documents = append(documents, *eventDocuments)
		}

		prevEventId = row.ID
	}

	// 上と重複
	// Time の処理
	if event != nil {
		event.Event.Datetimes = lo.Uniq[EventDatetime](event.Event.Datetimes)
		// Event の処理
		uniqDocuments := lo.Uniq[EventDocument](documents)
		event.Documents = &uniqDocuments

		events = append(events, *event)
	}

	return events, nil
}

func ConvEventList(queries *sqlc.Queries, eventName string, location string, setNotPublished bool, published bool) ([]EventEmbed, error) {
	eventsRows, err := queries.ListEvents(context.Background(), sqlc.ListEventsParams{
		SetEventName: eventName,
		SetLocation: sql.NullString{
			String: location,
			Valid:  true,
		},
		NotSetPublished: setNotPublished,
		SetPublished:    published,
	})
	if err != nil {
		return nil, err
	}
	if len(eventsRows) == 0 {
		return nil, nil
	}

	var prevEventId = eventsRows[0].ID
	var events []EventEmbed
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

		if prevEventId != row.ID {
			// Time の処理
			event.Event.Datetimes = lo.Uniq[EventDatetime](event.Event.Datetimes)

			events = append(events, *event)
			event = nil
		} else {
			// Time の処理
			event.Event.Datetimes = append(event.Event.Datetimes, *eventDatetime)
		}

		prevEventId = row.ID
	}

	// 上と重複
	// Time の処理
	if event != nil {
		event.Event.Datetimes = lo.Uniq[EventDatetime](event.Event.Datetimes)

		events = append(events, *event)
	}

	return events, nil
}
