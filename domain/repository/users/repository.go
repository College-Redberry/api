package users

import "com.redberry.api/domain/entities"

type Repository interface {
	Create(user entities.User) (int64, error)
	GetByEmail(email string) (entities.User, error)
	Update(userID int, updates entities.User) (int64, error)
	Delete(userID int) (int64, error)
}
