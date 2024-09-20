package messages

import (
	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/repository/messages"
)

type Service struct {
	messagesRepo *messages.Repository
}

func New() *Service {
	return &Service{
		messagesRepo: messages.New(),
	}
}

// Create cria uma nova mensagem.
func (service *Service) Create(message *entities.Message) (int, error) {
	return service.messagesRepo.Create(message)
}

// GetByID recupera uma mensagem pelo ID.
func (service *Service) GetByID(id int) (*entities.Message, error) {
	return service.messagesRepo.GetByID(id)
}

// Update atualiza uma mensagem existente.
func (service *Service) Update(message *entities.Message) error {
	return service.messagesRepo.Update(message)
}

// Delete remove uma mensagem pelo ID.
func (service *Service) Delete(id int) error {
	return service.messagesRepo.Delete(id)
}

// GetAll recupera todas as mensagens.
func (service *Service) GetAll() ([]*entities.Message, error) {
	return service.messagesRepo.GetAll()
}
