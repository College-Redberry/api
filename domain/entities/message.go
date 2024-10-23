package entities

import (
	"time"
)

type Message struct {
	ID              int       `json:"id" swaggerignore:"true"`
	CardID          int16     `json:"card_id"`
	CreatedAt       time.Time `json:"created_at" swaggerignore:"true" swaggerignore:"true"`
	UpdatedAt       time.Time `json:"updated_at,omitempty" swaggerignore:"true"`
	UserID          int16     `json:"user_id"`
	ParentMessageID int16     `json:"parent_message_id,omitempty"`
}
