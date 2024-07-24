package users

import (
	"context"
	"database/sql"

	"com.redberry.api/infrastructure/models"
	"com.redberry.api/infrastructure/postgres"
	"github.com/georgysavva/scany/sqlscan"
)

type Dao struct {
	db *sql.DB
}

func New() *Dao {
	return &Dao{
		db: postgres.Connect(),
	}
}

func (dao Dao) GetByEmail(email string) (*models.User, error) {
	var user models.User

	err := sqlscan.Get(context.Background(), dao.db, &user, `
		SELECT id, name, email, password, is_admin, profile_image FROM account.users WHERE email=$1
	`, email)

	return &user, err
}

func (dao Dao) Insert(user *models.User) (*models.User, error) {
	var insertedUser models.User

	err := sqlscan.Get(context.Background(), dao.db, &insertedUser, `
		INSERT INTO account."users"
		("name", email, "password", is_admin, profile_image)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id, name, email, is_admin, profile_image
	`, user.Name, user.Email, user.Password, user.IsAdmin, user.ProfileImage.String)

	return &insertedUser, err
}
