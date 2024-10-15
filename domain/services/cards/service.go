package cards

import (
	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/repository/cards"
	dao "com.redberry.api/infrastructure/dao/cards"
)

type Service struct {
	repository cards.Repository
}

func New() *Service {
	return &Service{
		repository: dao.New(),
	}
}

func (service *Service) Create(card entities.Card) (int64, error) {
	return service.repository.Create(card)
}

func (service *Service) GetByID(cardID int) (entities.Card, error) {
	return service.repository.GetByID(cardID)
}

func (service *Service) Update(cardID int, updates entities.Card) (int64, error) {
	return service.repository.Update(cardID, updates)
}

func (service *Service) Delete(cardID int) (int64, error) {
	return service.repository.Delete(cardID)
}
