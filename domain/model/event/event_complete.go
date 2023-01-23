package event

import (
	"github.com/jmoiron/sqlx"
	"prc_hub_back/domain/model/user"
)

func CompleteEvent(db *sqlx.DB, id int64, requestUser user.User) (Event, error) {
	completed := true
	return UpdateEvent(db, id, UpdateEventParam{Completed: &completed}, requestUser)
}
