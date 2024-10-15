package entities

import (
	"time"
)

type Card struct {
	ID                  int       `json:"id"`
	Title               string    `json:"title"`
	Description         string    `json:"description"`
	CreatedAt           time.Time `json:"created_at"`
	StartAt             time.Time `json:"start_at,omitempty"`
	UpdatedAt           time.Time `json:"updated_at,omitempty"`
	FinishedAt          time.Time `json:"finished_at,omitempty"`
	EstimatedFinishedAt time.Time `json:"estimated_finished_at,omitempty"`
	StatusID            int16     `json:"status_id"`
	ManagerID           int16     `json:"manager_id"`
	AssignedID          int16     `json:"assigned_id"`
	PriorityID          int16     `json:"priority_id"`
	ParentCardID        int16     `json:"parent_card_id,omitempty"`
}
