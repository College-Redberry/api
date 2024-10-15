package priorities

import "com.redberry.api/domain/entities"

type Repository interface {
	GetByID(priorityID int) (entities.Priority, error)
	Create(priority entities.Priority) (int64, error)
	Update(priorityID int, updates entities.Priority) (int64, error)
	Delete(priorityID int) (int64, error)
}
