package messages

import (
	"database/sql"

	"com.redberry.api/domain/entities"
	"com.redberry.api/infrastructure/dao/messages"
	"com.redberry.api/infrastructure/models"
)

type Repository struct {
	dao *messages.Dao
}

func New() *Repository {
	return &Repository{
		dao: messages.NewMessagesDao(),
	}
}

// Create insere uma nova mensagem no repositório.
func (repo *Repository) Create(message *entities.Message) (int, error) {
	createdMessage := &models.Message{
		CardID:          message.CardID,
		UserID:          sql.NullInt64{Int64: message.UserID},
		ParentMessageID: sql.NullInt64{Int64: message.ParentMessageID},
	}

	messageID, err := repo.dao.Create(createdMessage)
	if err != nil {
		return 0, err
	}

	return messageID, nil // Retorna o ID da mensagem criada
}

// GetByID recupera uma mensagem pelo ID.
func (repo *Repository) GetByID(id int) (*entities.Message, error) {
	messageModel, err := repo.dao.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &entities.Message{
		ID:              messageModel.ID,
		CardID:          messageModel.CardID,
		CreatedAt:       messageModel.CreatedAt,
		UpdatedAt:       messageModel.UpdatedAt,
		UserID:          messageModel.UserID.Int64,
		ParentMessageID: messageModel.ParentMessageID.Int64,
	}, nil
}

// Update atualiza uma mensagem existente no repositório.
func (repo *Repository) Update(message *entities.Message) error {
	messageModel := &models.Message{
		ID:              message.ID,
		CardID:          message.CardID,
		UserID:          sql.NullInt64{Int64: message.UserID},
		ParentMessageID: sql.NullInt64{Int64: message.ParentMessageID},
	}

	return repo.dao.Update(messageModel)
}

// Delete remove uma mensagem do repositório pelo ID.
func (repo *Repository) Delete(id int) error {
	return repo.dao.Delete(id)
}

// GetAll recupera todas as mensagens do repositório.
func (repo *Repository) GetAll() ([]*entities.Message, error) {
	messagesModel, err := repo.dao.GetAll()
	if err != nil {
		return nil, err
	}

	var messages []*entities.Message
	for _, m := range messagesModel {
		messages = append(messages, &entities.Message{
			ID:              m.ID,
			CardID:          m.CardID,
			CreatedAt:       m.CreatedAt,
			UpdatedAt:       m.UpdatedAt,
			UserID:          m.UserID.Int64,
			ParentMessageID: m.ParentMessageID.Int64,
		})
	}

	return messages, nil
}
