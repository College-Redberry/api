package messages

import "com.redberry.api/domain/entities"

type Repository interface {
	GetByID(cardID int) (entities.Message, error)
	Create(card entities.Message) (int64, error)
	Update(cardID int, updates entities.Message) (int64, error)
	Delete(cardID int) (int64, error)
}
