package entities

import (
	"time"
)

type Message struct {
	ID              int       `json:"id"`
	CardID          int16     `json:"card_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	UserID          int16     `json:"user_id"`
	ParentMessageID int16     `json:"parent_message_id,omitempty"`
}
