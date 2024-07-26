package users

import (
	"database/sql"

	"com.redberry.api/domain/entities"
	"com.redberry.api/infrastructure/dao/users"
	"com.redberry.api/infrastructure/models"
)

type Repository struct {
	dao *users.Dao
}

func New() *Repository {
	return &Repository{
		dao: users.New(),
	}
}

func (repo *Repository) GetByEmail(email string) (*entities.User, error) {
	user, err := repo.dao.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		IsAdmin:      user.IsAdmin,
		ProfileImage: user.ProfileImage.String,
	}, nil
}

func (repo *Repository) Create(user *entities.User) (*entities.User, error) {
	createdUser, err := repo.dao.Insert(&models.User{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		IsAdmin:      user.IsAdmin,
		ProfileImage: sql.NullString{String: user.ProfileImage},
	})
	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:           createdUser.ID,
		Name:         createdUser.Name,
		Email:        createdUser.Email,
		Password:     createdUser.Password,
		IsAdmin:      createdUser.IsAdmin,
		ProfileImage: createdUser.ProfileImage.String,
	}, nil
}
