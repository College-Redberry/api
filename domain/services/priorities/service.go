package priorities

import (
	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/repository/priorities"
	dao "com.redberry.api/infrastructure/dao/priorities"
)

type Service struct {
	repository priorities.Repository
}

func New() *Service {
	return &Service{
		repository: dao.New(),
	}
}

func (service *Service) Create(priority entities.Priority) (int64, error) {
	return service.repository.Create(priority)
}

func (service *Service) GetByID(priorityID int) (entities.Priority, error) {
	return service.repository.GetByID(priorityID)
}

func (service *Service) Update(priorityID int, updates entities.Priority) (int64, error) {
	return service.repository.Update(priorityID, updates)
}

func (service *Service) Delete(priorityID int) (int64, error) {
	return service.repository.Delete(priorityID)
}
