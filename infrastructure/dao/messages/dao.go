package dao

import (
	"context"
	"database/sql"

	"com.redberry.api/domain/entities"
	"com.redberry.api/infrastructure/models"
	"com.redberry.api/infrastructure/postgres"
	"github.com/georgysavva/scany/sqlscan"
)

type Dao struct {
	db *sql.DB
}

func New() *Dao {
	return &Dao{
		db: postgres.Connect(),
	}
}

func (dao *Dao) GetByID(messageID int) (entities.Message, error) {
	var message models.Message

	err := sqlscan.Get(context.Background(), dao.db, &message, `
        SELECT 
            id,
            card_id,
            created_at,
            updated_at,
            user_id,
            parent_message_id
        FROM communication.messages WHERE id=$1
    `, messageID)
	if err != nil {
		return entities.Message{}, err
	}

	return entities.Message{
		ID:              message.ID,
		CardID:          message.CardID,
		CreatedAt:       message.CreatedAt,
		UpdatedAt:       message.UpdatedAt.Time,
		UserID:          message.UserID,
		ParentMessageID: int16(message.ParentMessageID.Int64),
	}, nil
}

func (dao *Dao) Create(message entities.Message) (int64, error) {
	var id int64

	var parentMessageID *int16
	if message.ParentMessageID != 0 {
		parentMessageID = &message.ParentMessageID
	}

	err := sqlscan.Get(context.Background(), dao.db, &id, `        
		INSERT INTO communication.messages (
            card_id,
            user_id,
            parent_message_id
        ) VALUES ($1, $2, $3)
        RETURNING id
    `,
		message.CardID,
		message.UserID,
		parentMessageID,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *Dao) Update(messageID int, updates entities.Message) (int64, error) {
	var id int64

	err := sqlscan.Get(context.Background(), dao.db, &id, `        
		UPDATE communication.messages
        SET 
            card_id = COALESCE($2, card_id),
            updated_at = NOW(),
            user_id = COALESCE($3, user_id)
        WHERE id = $1 RETURNING id
    `,
		messageID,
		updates.CardID,
		updates.UserID,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *Dao) Delete(messageID int) (int64, error) {
	var id int64

	err := sqlscan.Get(context.Background(), dao.db, &id, `        
		DELETE FROM communication.messages WHERE id = $1 RETURNING id
    `, messageID)
	if err != nil {
		return 0, err
	}

	return id, nil
}
