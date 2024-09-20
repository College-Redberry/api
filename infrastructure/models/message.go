package models

import (
	"database/sql"
	"time"
)

type Message struct {
	ID              int           `db:"id"`
	CardID          int           `db:"card_id"`
	CreatedAt       time.Time     `db:"created_at"`
	UpdatedAt       time.Time     `db:"updated_at"`
	UserID          sql.NullInt64 `db:"user_id"`
	ParentMessageID sql.NullInt64 `db:"parent_message_id"`
}
