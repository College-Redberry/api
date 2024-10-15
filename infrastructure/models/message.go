package models

import (
	"database/sql"
	"time"
)

type Message struct {
	ID              int           `db:"id"`
	CardID          int16         `db:"card_id"`
	CreatedAt       time.Time     `db:"created_at"`
	UpdatedAt       sql.NullTime  `db:"updated_at"`
	UserID          int16         `db:"user_id"`
	ParentMessageID sql.NullInt64 `db:"parent_message_id"`
}
