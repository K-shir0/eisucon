// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package sqlc

import (
	"database/sql"
	"time"
)

type Document struct {
	ID      int32
	EventID int32
	Name    string
	Url     string
}

type Event struct {
	ID          int32
	Name        string
	Description sql.NullString
	Location    sql.NullString
	Published   bool
	Completed   bool
	UserID      int32
}

type EventDatetime struct {
	EventID int32
	Start   time.Time
	End     time.Time
}

type User struct {
	ID                  int32
	Name                string
	Email               string
	Password            string
	PostEventAvailabled bool
	Manage              bool
	Admin               bool
	TwitterID           sql.NullString
	GithubUsername      sql.NullString
}

type UserStar struct {
	ID           int32
	TargetUserID int32
}
