package boards

import "com.redberry.api/domain/entities"

type Repository interface {
	GetByID(boardID int) (entities.Board, error)
	Create(board entities.Board) (int64, error)
	Update(boardID int, updates entities.Board) (int64, error)
	Delete(boardID int) (int64, error)
}
