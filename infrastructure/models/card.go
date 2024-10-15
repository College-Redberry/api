package models

import (
	"database/sql"
	"time"
)

type Card struct {
	ID                  int           `db:"id"`
	Title               string        `db:"title"`
	Description         string        `db:"description"`
	CreatedAt           time.Time     `db:"created_at"`
	StartAt             sql.NullTime  `db:"start_at"`
	UpdatedAt           sql.NullTime  `db:"updated_at"`
	FinishedAt          sql.NullTime  `db:"finished_at"`
	EstimatedFinishedAt sql.NullTime  `db:"estimated_finished_at"`
	StatusID            int16         `db:"status_id"`
	ManagerID           int16         `db:"manager_id"`
	AssignedID          int16         `db:"assigned_id"`
	PriorityID          int16         `db:"priority_id"`
	ParentCardID        sql.NullInt64 `db:"parent_card_id"`
}
