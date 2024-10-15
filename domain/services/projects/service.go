package projects

import (
	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/repository/projects"
	dao "com.redberry.api/infrastructure/dao/project"
)

type Service struct {
	repository projects.Repository
}

func New() *Service {
	return &Service{
		repository: dao.New(),
	}
}

func (service *Service) Create(project entities.Project) (int64, error) {
	return service.repository.Create(project)
}

func (service *Service) GetByID(projectID int) (entities.Project, error) {
	return service.repository.GetByID(projectID)
}

func (service *Service) Update(projectID int, updates entities.Project) (int64, error) {
	return service.repository.Update(projectID, updates)
}

func (service *Service) Delete(projectID int) (int64, error) {
	return service.repository.Delete(projectID)
}
