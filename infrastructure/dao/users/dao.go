package users

import (
	"context"
	"database/sql"
	"time"

	"com.redberry.api/domain/entities"
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

func (dao Dao) GetByEmail(email string) (entities.User, error) {
	var user models.User

	err := sqlscan.Get(context.Background(), dao.db, &user, `
		SELECT id, name, email, password, is_admin, profile_image FROM account.users WHERE email=$1
	`, email)
	if err != nil {
		return entities.User{}, err
	}

	return entities.User{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		IsAdmin:      user.IsAdmin,
		ProfileImage: user.ProfileImage.String,
	}, nil
}

func (dao Dao) Create(user entities.User) (int64, error) {
	result, err := dao.db.ExecContext(context.Background(), `
		INSERT INTO account."users"
		("name", email, "password", is_admin, profile_image)
		VALUES($1, $2, $3, $4, $5)
	`, user.Name, user.Email, user.Password, user.IsAdmin, user.ProfileImage)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
func (dao *Dao) Update(userID int, updates entities.User) (int64, error) {
	result, err := dao.db.ExecContext(context.Background(), `
        UPDATE account.users 
        SET 
            name = COALESCE($2, name),
            profile_image = COALESCE($3, profile_image),
            updated_at = $4
        WHERE id = $1
    `,
		userID,
		updates.Name,
		updates.ProfileImage,
		time.Now(),
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (dao *Dao) Delete(userD int) (int64, error) {
	result, err := dao.db.ExecContext(context.Background(), `
        DELETE FROM account.users WHERE id = $1
    `, userD)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
