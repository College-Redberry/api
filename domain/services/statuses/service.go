package statuses

import (
	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/repository/statuses"
	dao "com.redberry.api/infrastructure/dao/statuses"
)

type Service struct {
	repository statuses.Repository
}

func New() *Service {
	return &Service{
		repository: dao.New(),
	}
}

func (service *Service) Create(status entities.Status) (int64, error) {
	return service.repository.Create(status)
}

func (service *Service) GetByID(statusID int) (entities.Status, error) {
	return service.repository.GetByID(statusID)
}

func (service *Service) Update(statusID int, updates entities.Status) (int64, error) {
	return service.repository.Update(statusID, updates)
}

func (service *Service) Delete(statusID int) (int64, error) {
	return service.repository.Delete(statusID)
}
