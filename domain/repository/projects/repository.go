package projects

import "com.redberry.api/domain/entities"

type Repository interface {
	GetByID(projectID int) (entities.Project, error)
	Create(project entities.Project) (int64, error)
	Update(projectID int, updates entities.Project) (int64, error)
	Delete(projectID int) (int64, error)
}
