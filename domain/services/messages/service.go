package messages

import (
	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/repository/messages"
	dao "com.redberry.api/infrastructure/dao/messages"
)

type Service struct {
	repository messages.Repository
}

func New() *Service {
	return &Service{
		repository: dao.New(),
	}
}

func (service *Service) Create(message entities.Message) (int64, error) {
	return service.repository.Create(message)
}

func (service *Service) GetByID(messageID int) (entities.Message, error) {
	return service.repository.GetByID(messageID)
}

func (service *Service) Update(messageID int, updates entities.Message) (int64, error) {
	return service.repository.Update(messageID, updates)
}

func (service *Service) Delete(messageID int) (int64, error) {
	return service.repository.Delete(messageID)
}
