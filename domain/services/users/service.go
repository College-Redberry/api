package users

import (
	"com.redberry.api/application/constants"
	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/repository/users"
	dao "com.redberry.api/infrastructure/dao/users"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository users.Repository
}

func New() *Service {
	return &Service{
		repository: dao.New(),
	}
}

func (service *Service) Register(user entities.User) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), constants.PasswordHashCost)
	if err != nil {
		return 0, err
	}

	user.Password = string(hashedPassword)

	return service.repository.Create(user)
}

func (service *Service) GetByEmail(email string) (entities.User, error) {
	return service.repository.GetByEmail(email)
}

func (service *Service) Update(userID int, updates entities.User) (int64, error) {
	return service.repository.Update(userID, updates)
}

func (service *Service) Delete(userID int) (int64, error) {
	return service.repository.Delete(userID)
}
