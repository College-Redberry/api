package auth

import (
	"com.redberry.api/domain/repository/users"
	dao "com.redberry.api/infrastructure/dao/users"
	"com.redberry.api/infrastructure/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	usersRepo users.Repository
}

func New() *Service {
	return &Service{
		usersRepo: dao.New(),
	}
}

func (service *Service) Login(email, password string) (string, error) {
	user, err := service.usersRepo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	return jwt.Generate(user.IsAdmin)
}
