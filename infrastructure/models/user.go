package models

import (
	"database/sql"
)

type User struct {
	ID           int            `db:"id"`
	Name         string         `db:"name"`
	Email        string         `db:"email"`
	Password     string         `db:"password,omitempty"`
	IsAdmin      bool           `db:"is_admin"`
	ProfileImage sql.NullString `db:"profile_image"`
}
