package users

import (
	"com.redberry.api/application/constants"
	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/repository/users"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	usersRepo *users.Repository
}

func New() *Service {
	return &Service{
		usersRepo: users.New(),
	}
}

func (service *Service) Register(user *entities.User) (*entities.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), constants.PasswordHashCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	return service.usersRepo.Create(user)
}
