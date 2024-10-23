package entities

import "time"

type Project struct {
	ID          int       `json:"id" swaggerignore:"true"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at" swaggerignore:"true"`
	UpdatedAt   time.Time `json:"updated_at" swaggerignore:"true"`
}
