package entities

import "time"

type Board struct {
	ID          int       `json:"id" swaggerignore:"true"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ManagerID   int16     `json:"manager_id"`
	CreatedAt   time.Time `json:"created_at" swaggerignore:"true"`
	UpdatedAt   time.Time `json:"updated_at" swaggerignore:"true"`
	ProjectID   int16     `json:"project_id"`
}
