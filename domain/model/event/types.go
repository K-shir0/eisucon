package event

import (
	"prc_hub_back/domain/model/user"
	"time"
)

type Event struct {
	Id          int64           `json:"id"`
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	Location    *string         `json:"location,omitempty"`
	Datetimes   []EventDatetime `db:"event_datetimes" json:"datetimes"`
	Published   bool            `json:"published"`
	Completed   bool            `json:"completed"`
	UserId      int64           `json:"user_id"`
}

type EventDatetime struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end" dh:"end"`
}

type EventDocument struct {
	EventId int64  `json:"event_id"`
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
}

type EventEmbed struct {
	Event     `db:"events"`
	User      *user.User       `db:"users" json:"user,omitempty"`
	Documents *[]EventDocument `db:"documents" json:"documents,omitempty"`
}
