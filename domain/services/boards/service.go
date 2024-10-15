package boards

import (
	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/repository/boards"
	dao "com.redberry.api/infrastructure/dao/boards"
)

type Service struct {
	repository boards.Repository
}

func New() *Service {
	return &Service{
		repository: dao.New(),
	}
}

func (service *Service) Create(board entities.Board) (int64, error) {
	return service.repository.Create(board)
}

func (service *Service) GetByID(boardID int) (entities.Board, error) {
	return service.repository.GetByID(boardID)
}

func (service *Service) Update(boardID int, updates entities.Board) (int64, error) {
	return service.repository.Update(boardID, updates)
}

func (service *Service) Delete(boardID int) (int64, error) {
	return service.repository.Delete(boardID)
}
