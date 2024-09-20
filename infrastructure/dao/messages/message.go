package messages

import (
	"context"
	"database/sql"

	"com.redberry.api/infrastructure/models"
	"com.redberry.api/infrastructure/postgres"
	"github.com/georgysavva/scany/sqlscan"
)

type Dao struct {
	db *sql.DB
}

func NewMessagesDao() *Dao {
	return &Dao{
		db: postgres.Connect(),
	}
}

// Create insere uma nova mensagem no banco de dados e retorna o ID da mensagem.
func (dao *Dao) Create(message *models.Message) (int, error) {
	query := `
		INSERT INTO communication.messages (card_id, created_at, updated_at, user_id, parent_message_id) 
		VALUES ($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $2, $3) 
		RETURNING id`

	err := dao.db.QueryRowContext(context.Background(), query, message.CardID, message.UserID, message.ParentMessageID).Scan(&message.ID)
	if err != nil {
		return 0, err // Retorna 0 se houver um erro
	}
	return message.ID, nil // Retorna o ID da mensagem criada
}

// GetByID recupera uma mensagem pelo ID.
func (dao *Dao) GetByID(id int) (*models.Message, error) {
	var message models.Message

	err := sqlscan.Get(context.Background(), dao.db, &message, `
		SELECT id, card_id, created_at, updated_at, user_id, parent_message_id 
		FROM communication.messages WHERE id=$1
	`, id)

	if err != nil {
		return nil, err
	}

	return &message, nil
}

// Update atualiza uma mensagem existente.
func (dao *Dao) Update(message *models.Message) error {
	query := `
		UPDATE communication.messages 
		SET card_id = $1, updated_at = CURRENT_TIMESTAMP, user_id = $2, parent_message_id = $3 
		WHERE id = $4`

	_, err := dao.db.ExecContext(context.Background(), query, message.CardID, message.UserID, message.ParentMessageID, message.ID)
	return err
}

// Delete remove uma mensagem do banco de dados pelo ID.
func (dao *Dao) Delete(id int) error {
	query := `
		DELETE FROM communication.messages WHERE id = $1`

	_, err := dao.db.ExecContext(context.Background(), query, id)
	return err
}

// GetAll recupera todas as mensagens do banco de dados.
func (dao *Dao) GetAll() ([]models.Message, error) {
	var messages []models.Message

	err := sqlscan.Select(context.Background(), dao.db, &messages, `
		SELECT id, card_id, created_at, updated_at, user_id, parent_message_id 
		FROM communication.messages`)

	if err != nil {
		return nil, err
	}

	return messages, nil
}
