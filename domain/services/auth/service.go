package auth

import (
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

func (service *Service) Login(email, password string) error {
	user, err := service.usersRepo.GetByEmail(email)
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
