package statuses

import "com.redberry.api/domain/entities"

type Repository interface {
	GetByID(statusID int) (entities.Status, error)
	Create(status entities.Status) (int64, error)
	Update(statusID int, updates entities.Status) (int64, error)
	Delete(statusID int) (int64, error)
}
