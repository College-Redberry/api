package entities

import "time"

type Message struct {
	ID              int       `json:"id"`
	CardID          int       `json:"card_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	UserID          int64     `json:"user_id"`
	ParentMessageID int64     `json:"parent_message_id"`
}
