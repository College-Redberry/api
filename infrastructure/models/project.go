package models

import (
	"database/sql"
	"time"
)

type Project struct {
	ID          int          `db:"id"`
	Name        string       `db:"name"`
	Description string       `db:"description"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}
