package models

import (
	"database/sql"
	"time"
)

type Board struct {
	ID          int          `db:"id"`
	Name        string       `db:"name"`
	Description string       `db:"description"`
	ManagerID   int16        `db:"manager_id"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
	ProjectID   int16        `db:"project_id"`
}
