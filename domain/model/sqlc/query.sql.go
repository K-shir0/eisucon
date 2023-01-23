// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const getEvent = `-- name: GetEvent :many
SELECT events.id,
       events.name,
       events.description,
       events.location,
       events.published,
       events.completed,
       events.user_id,
       event_datetimes.event_id,
       event_datetimes.start,
       event_datetimes.end
FROM events
         JOIN event_datetimes ON events.id = event_datetimes.event_id
GROUP BY events.id, events.name, events.description, events.location, events.published, events.completed,
         events.user_id, event_datetimes.event_id, event_datetimes.start, event_datetimes.end
HAVING events.id = ?
`

type GetEventRow struct {
	ID          int32
	Name        string
	Description sql.NullString
	Location    sql.NullString
	Published   bool
	Completed   bool
	UserID      int32
	EventID     int32
	Start       time.Time
	End         time.Time
}

func (q *Queries) GetEvent(ctx context.Context, setEventID int32) ([]GetEventRow, error) {
	rows, err := q.db.QueryContext(ctx, getEvent, setEventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetEventRow
	for rows.Next() {
		var i GetEventRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Location,
			&i.Published,
			&i.Completed,
			&i.UserID,
			&i.EventID,
			&i.Start,
			&i.End,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEventWithDocuments = `-- name: GetEventWithDocuments :many
SELECT events.id,
       events.name,
       events.description,
       events.location,
       events.published,
       events.completed,
       events.user_id,
       event_datetimes.event_id,
       event_datetimes.start,
       event_datetimes.end,
       documents.id       AS document_id,
       documents.event_id AS document_event_id,
       documents.name     AS document_name,
       documents.url
FROM events
         JOIN event_datetimes ON events.id = event_datetimes.event_id
         JOIN documents ON events.id = documents.event_id
GROUP BY events.id, events.name, events.description, events.location, events.published, events.completed,
         events.user_id, event_datetimes.event_id, event_datetimes.start, event_datetimes.end, documents.id,
         documents.event_id, documents.name, documents.url
HAVING events.id = ?
`

type GetEventWithDocumentsRow struct {
	ID              int32
	Name            string
	Description     sql.NullString
	Location        sql.NullString
	Published       bool
	Completed       bool
	UserID          int32
	EventID         int32
	Start           time.Time
	End             time.Time
	DocumentID      int32
	DocumentEventID int32
	DocumentName    string
	Url             string
}

func (q *Queries) GetEventWithDocuments(ctx context.Context, setEventID int32) ([]GetEventWithDocumentsRow, error) {
	rows, err := q.db.QueryContext(ctx, getEventWithDocuments, setEventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetEventWithDocumentsRow
	for rows.Next() {
		var i GetEventWithDocumentsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Location,
			&i.Published,
			&i.Completed,
			&i.UserID,
			&i.EventID,
			&i.Start,
			&i.End,
			&i.DocumentID,
			&i.DocumentEventID,
			&i.DocumentName,
			&i.Url,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEventWithUser = `-- name: GetEventWithUser :many
SELECT events.id,
       events.name,
       events.description,
       events.location,
       events.published,
       events.completed,
       events.user_id,
       event_datetimes.event_id,
       event_datetimes.start,
       event_datetimes.end,
       users.id,
       users.name,
       users.email,
       users.password,
       users.post_event_availabled,
       users.manage,
       users.admin,
       users.twitter_id,
       users.github_username,
       COUNT(user_stars.target_user_id) as star_count
FROM events
         JOIN event_datetimes ON events.id = event_datetimes.event_id
         JOIN users ON events.user_id = users.id
         LEFT JOIN user_stars ON users.id = user_stars.target_user_id
GROUP BY events.id, events.name, events.description, events.location, events.published, events.completed,
         events.user_id, event_datetimes.event_id, event_datetimes.start, event_datetimes.end, users.id, users.name,
         users.email, users.password, users.post_event_availabled, users.manage, users.admin, users.twitter_id,
         users.github_username
HAVING events.id = ?
`

type GetEventWithUserRow struct {
	ID                  int32
	Name                string
	Description         sql.NullString
	Location            sql.NullString
	Published           bool
	Completed           bool
	UserID              int32
	EventID             int32
	Start               time.Time
	End                 time.Time
	ID_2                int32
	Name_2              string
	Email               string
	Password            string
	PostEventAvailabled bool
	Manage              bool
	Admin               bool
	TwitterID           sql.NullString
	GithubUsername      sql.NullString
	StarCount           int64
}

func (q *Queries) GetEventWithUser(ctx context.Context, setEventID int32) ([]GetEventWithUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getEventWithUser, setEventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetEventWithUserRow
	for rows.Next() {
		var i GetEventWithUserRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Location,
			&i.Published,
			&i.Completed,
			&i.UserID,
			&i.EventID,
			&i.Start,
			&i.End,
			&i.ID_2,
			&i.Name_2,
			&i.Email,
			&i.Password,
			&i.PostEventAvailabled,
			&i.Manage,
			&i.Admin,
			&i.TwitterID,
			&i.GithubUsername,
			&i.StarCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEventWithUserAndDocuments = `-- name: GetEventWithUserAndDocuments :many
SELECT events.id,
       events.name,
       events.description,
       events.location,
       events.published,
       events.completed,
       events.user_id,
       event_datetimes.event_id,
       event_datetimes.start,
       event_datetimes.end,
       documents.id                     AS document_id,
       documents.event_id               AS document_event_id,
       documents.name                   AS document_name,
       documents.url,
       users.id,
       users.name,
       users.email,
       users.password,
       users.post_event_availabled,
       users.manage,
       users.admin,
       users.twitter_id,
       users.github_username,
       COUNT(user_stars.target_user_id) as star_count
FROM events
         JOIN event_datetimes ON events.id = event_datetimes.event_id
         JOIN documents ON events.id = documents.event_id
         JOIN users ON events.user_id = users.id
         LEFT JOIN user_stars ON users.id = user_stars.target_user_id
GROUP BY events.id, events.name, events.description, events.location, events.published, events.completed,
         events.user_id, event_datetimes.event_id, event_datetimes.start, event_datetimes.end, documents.id,
         documents.event_id, documents.name, documents.url, users.id, users.name, users.email, users.password,
         users.post_event_availabled, users.manage, users.admin, users.twitter_id, users.github_username
HAVING events.id = ?
`

type GetEventWithUserAndDocumentsRow struct {
	ID                  int32
	Name                string
	Description         sql.NullString
	Location            sql.NullString
	Published           bool
	Completed           bool
	UserID              int32
	EventID             int32
	Start               time.Time
	End                 time.Time
	DocumentID          int32
	DocumentEventID     int32
	DocumentName        string
	Url                 string
	ID_2                int32
	Name_2              string
	Email               string
	Password            string
	PostEventAvailabled bool
	Manage              bool
	Admin               bool
	TwitterID           sql.NullString
	GithubUsername      sql.NullString
	StarCount           int64
}

func (q *Queries) GetEventWithUserAndDocuments(ctx context.Context, setEventID int32) ([]GetEventWithUserAndDocumentsRow, error) {
	rows, err := q.db.QueryContext(ctx, getEventWithUserAndDocuments, setEventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetEventWithUserAndDocumentsRow
	for rows.Next() {
		var i GetEventWithUserAndDocumentsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Location,
			&i.Published,
			&i.Completed,
			&i.UserID,
			&i.EventID,
			&i.Start,
			&i.End,
			&i.DocumentID,
			&i.DocumentEventID,
			&i.DocumentName,
			&i.Url,
			&i.ID_2,
			&i.Name_2,
			&i.Email,
			&i.Password,
			&i.PostEventAvailabled,
			&i.Manage,
			&i.Admin,
			&i.TwitterID,
			&i.GithubUsername,
			&i.StarCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT u.id,
       u.name,
       u.email,
       u.password,
       u.post_event_availabled,
       u.manage,
       u.admin,
       u.twitter_id,
       u.github_username,
       COUNT(s.target_user_id) AS star_count
FROM users u
         LEFT JOIN user_stars s ON u.id = s.target_user_id
GROUP BY u.id
HAVING u.email LIKE CASE
                        WHEN ? != '%'
                            THEN ?
                        ELSE u.email
    END
`

type GetUserParams struct {
	SetEmail string
}

type GetUserRow struct {
	ID                  int32
	Name                string
	Email               string
	Password            string
	PostEventAvailabled bool
	Manage              bool
	Admin               bool
	TwitterID           sql.NullString
	GithubUsername      sql.NullString
	StarCount           int64
}

func (q *Queries) GetUser(ctx context.Context, arg GetUserParams) (GetUserRow, error) {
	row := q.db.QueryRowContext(ctx, getUser, arg.SetEmail, arg.SetEmail)
	var i GetUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.PostEventAvailabled,
		&i.Manage,
		&i.Admin,
		&i.TwitterID,
		&i.GithubUsername,
		&i.StarCount,
	)
	return i, err
}

const listEvents = `-- name: ListEvents :many
SELECT events.id,
       events.name,
       events.description,
       events.location,
       events.published,
       events.completed,
       events.user_id,
       event_datetimes.event_id,
       event_datetimes.start,
       event_datetimes.end
FROM events
         JOIN event_datetimes ON events.id = event_datetimes.event_id
GROUP BY events.id, events.name, events.description, events.location, events.published, events.completed,
         events.user_id, event_datetimes.event_id, event_datetimes.start, event_datetimes.end
HAVING events.name LIKE CASE
                            WHEN ? != '%'
                                THEN ?
                            ELSE events.name
    END
   AND events.location LIKE CASE
                                WHEN ? != '%'
                                    THEN ?
                                ELSE events.location
    END
   AND events.published = CASE
                              WHEN ? = false
                                  THEN ?
                              ELSE events.published
    END
`

type ListEventsParams struct {
	SetEventName    string
	SetLocation     sql.NullString
	NotSetPublished interface{}
	SetPublished    bool
}

type ListEventsRow struct {
	ID          int32
	Name        string
	Description sql.NullString
	Location    sql.NullString
	Published   bool
	Completed   bool
	UserID      int32
	EventID     int32
	Start       time.Time
	End         time.Time
}

func (q *Queries) ListEvents(ctx context.Context, arg ListEventsParams) ([]ListEventsRow, error) {
	rows, err := q.db.QueryContext(ctx, listEvents,
		arg.SetEventName,
		arg.SetEventName,
		arg.SetLocation,
		arg.SetLocation,
		arg.NotSetPublished,
		arg.SetPublished,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListEventsRow
	for rows.Next() {
		var i ListEventsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Location,
			&i.Published,
			&i.Completed,
			&i.UserID,
			&i.EventID,
			&i.Start,
			&i.End,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listEventsWithDocuments = `-- name: ListEventsWithDocuments :many
SELECT events.id,
       events.name,
       events.description,
       events.location,
       events.published,
       events.completed,
       events.user_id,
       event_datetimes.event_id,
       event_datetimes.start,
       event_datetimes.end,
       documents.id       AS document_id,
       documents.event_id AS document_event_id,
       documents.name     AS document_name,
       documents.url
FROM events
         JOIN event_datetimes ON events.id = event_datetimes.event_id
         JOIN documents ON events.id = documents.event_id
GROUP BY events.id, events.name, events.description, events.location, events.published, events.completed,
         events.user_id, event_datetimes.event_id, event_datetimes.start, event_datetimes.end, documents.id,
         documents.event_id, documents.name, documents.url
HAVING events.name LIKE CASE
                            WHEN ? != '%'
                                THEN ?
                            ELSE events.name
    END
   AND events.location LIKE CASE
                                WHEN ? != '%'
                                    THEN ?
                                ELSE events.location
    END
   AND events.published = CASE
                              WHEN ? = false
                                  THEN ?
                              ELSE events.published
    END
`

type ListEventsWithDocumentsParams struct {
	SetEventName    string
	SetLocation     sql.NullString
	NotSetPublished interface{}
	SetPublished    bool
}

type ListEventsWithDocumentsRow struct {
	ID              int32
	Name            string
	Description     sql.NullString
	Location        sql.NullString
	Published       bool
	Completed       bool
	UserID          int32
	EventID         int32
	Start           time.Time
	End             time.Time
	DocumentID      int32
	DocumentEventID int32
	DocumentName    string
	Url             string
}

func (q *Queries) ListEventsWithDocuments(ctx context.Context, arg ListEventsWithDocumentsParams) ([]ListEventsWithDocumentsRow, error) {
	rows, err := q.db.QueryContext(ctx, listEventsWithDocuments,
		arg.SetEventName,
		arg.SetEventName,
		arg.SetLocation,
		arg.SetLocation,
		arg.NotSetPublished,
		arg.SetPublished,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListEventsWithDocumentsRow
	for rows.Next() {
		var i ListEventsWithDocumentsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Location,
			&i.Published,
			&i.Completed,
			&i.UserID,
			&i.EventID,
			&i.Start,
			&i.End,
			&i.DocumentID,
			&i.DocumentEventID,
			&i.DocumentName,
			&i.Url,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listEventsWithUser = `-- name: ListEventsWithUser :many
SELECT events.id,
       events.name,
       events.description,
       events.location,
       events.published,
       events.completed,
       events.user_id,
       event_datetimes.event_id,
       event_datetimes.start,
       event_datetimes.end,
       users.id,
       users.name,
       users.email,
       users.password,
       users.post_event_availabled,
       users.manage,
       users.admin,
       users.twitter_id,
       users.github_username,
       COUNT(user_stars.target_user_id) as star_count
FROM events
         JOIN event_datetimes ON events.id = event_datetimes.event_id
         JOIN users ON events.user_id = users.id
         LEFT JOIN user_stars ON users.id = user_stars.target_user_id
GROUP BY events.id, events.name, events.description, events.location, events.published, events.completed,
         events.user_id, event_datetimes.event_id, event_datetimes.start, event_datetimes.end, users.id, users.name,
         users.email, users.password, users.post_event_availabled, users.manage, users.admin, users.twitter_id,
         users.github_username
HAVING events.name LIKE CASE
                            WHEN ? != '%'
                                THEN ?
                            ELSE events.name
    END
   AND events.location LIKE CASE
                                WHEN ? != '%'
                                    THEN ?
                                ELSE events.location
    END
   AND events.published = CASE
                              WHEN ? = false
                                  THEN ?
                              ELSE events.published
    END
`

type ListEventsWithUserParams struct {
	SetEventName    string
	SetLocation     sql.NullString
	NotSetPublished interface{}
	SetPublished    bool
}

type ListEventsWithUserRow struct {
	ID                  int32
	Name                string
	Description         sql.NullString
	Location            sql.NullString
	Published           bool
	Completed           bool
	UserID              int32
	EventID             int32
	Start               time.Time
	End                 time.Time
	ID_2                int32
	Name_2              string
	Email               string
	Password            string
	PostEventAvailabled bool
	Manage              bool
	Admin               bool
	TwitterID           sql.NullString
	GithubUsername      sql.NullString
	StarCount           int64
}

func (q *Queries) ListEventsWithUser(ctx context.Context, arg ListEventsWithUserParams) ([]ListEventsWithUserRow, error) {
	rows, err := q.db.QueryContext(ctx, listEventsWithUser,
		arg.SetEventName,
		arg.SetEventName,
		arg.SetLocation,
		arg.SetLocation,
		arg.NotSetPublished,
		arg.SetPublished,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListEventsWithUserRow
	for rows.Next() {
		var i ListEventsWithUserRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Location,
			&i.Published,
			&i.Completed,
			&i.UserID,
			&i.EventID,
			&i.Start,
			&i.End,
			&i.ID_2,
			&i.Name_2,
			&i.Email,
			&i.Password,
			&i.PostEventAvailabled,
			&i.Manage,
			&i.Admin,
			&i.TwitterID,
			&i.GithubUsername,
			&i.StarCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listEventsWithUserAndDocuments = `-- name: ListEventsWithUserAndDocuments :many
SELECT events.id,
       events.name,
       events.description,
       events.location,
       events.published,
       events.completed,
       events.user_id,
       event_datetimes.event_id,
       event_datetimes.start,
       event_datetimes.end,
       documents.id                     AS document_id,
       documents.event_id               AS document_event_id,
       documents.name                   AS document_name,
       documents.url,
       users.id,
       users.name,
       users.email,
       users.password,
       users.post_event_availabled,
       users.manage,
       users.admin,
       users.twitter_id,
       users.github_username,
       COUNT(user_stars.target_user_id) as star_count
FROM events
         JOIN event_datetimes ON events.id = event_datetimes.event_id
         JOIN documents ON events.id = documents.event_id
         JOIN users ON events.user_id = users.id
         LEFT JOIN user_stars ON users.id = user_stars.target_user_id
GROUP BY events.id, events.name, events.description, events.location, events.published, events.completed,
         events.user_id, event_datetimes.event_id, event_datetimes.start, event_datetimes.end, documents.id,
         documents.event_id, documents.name, documents.url, users.id, users.name, users.email, users.password,
         users.post_event_availabled, users.manage, users.admin, users.twitter_id, users.github_username
HAVING events.name LIKE CASE
                            WHEN ? != '%'
                                THEN ?
                            ELSE events.name
    END
   AND events.location LIKE CASE
                                WHEN ? != '%'
                                    THEN ?
                                ELSE events.location
    END
   AND events.published = CASE
                              WHEN ? = false
                                  THEN ?
                              ELSE events.published
    END
`

type ListEventsWithUserAndDocumentsParams struct {
	SetEventName    string
	SetLocation     sql.NullString
	NotSetPublished interface{}
	SetPublished    bool
}

type ListEventsWithUserAndDocumentsRow struct {
	ID                  int32
	Name                string
	Description         sql.NullString
	Location            sql.NullString
	Published           bool
	Completed           bool
	UserID              int32
	EventID             int32
	Start               time.Time
	End                 time.Time
	DocumentID          int32
	DocumentEventID     int32
	DocumentName        string
	Url                 string
	ID_2                int32
	Name_2              string
	Email               string
	Password            string
	PostEventAvailabled bool
	Manage              bool
	Admin               bool
	TwitterID           sql.NullString
	GithubUsername      sql.NullString
	StarCount           int64
}

func (q *Queries) ListEventsWithUserAndDocuments(ctx context.Context, arg ListEventsWithUserAndDocumentsParams) ([]ListEventsWithUserAndDocumentsRow, error) {
	rows, err := q.db.QueryContext(ctx, listEventsWithUserAndDocuments,
		arg.SetEventName,
		arg.SetEventName,
		arg.SetLocation,
		arg.SetLocation,
		arg.NotSetPublished,
		arg.SetPublished,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListEventsWithUserAndDocumentsRow
	for rows.Next() {
		var i ListEventsWithUserAndDocumentsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Location,
			&i.Published,
			&i.Completed,
			&i.UserID,
			&i.EventID,
			&i.Start,
			&i.End,
			&i.DocumentID,
			&i.DocumentEventID,
			&i.DocumentName,
			&i.Url,
			&i.ID_2,
			&i.Name_2,
			&i.Email,
			&i.Password,
			&i.PostEventAvailabled,
			&i.Manage,
			&i.Admin,
			&i.TwitterID,
			&i.GithubUsername,
			&i.StarCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
