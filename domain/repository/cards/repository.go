package cards

import "com.redberry.api/domain/entities"

type Repository interface {
	GetByID(cardID int) (entities.Card, error)
	Create(card entities.Card) (int64, error)
	Update(cardID int, updates entities.Card) (int64, error)
	Delete(cardID int) (int64, error)
}
