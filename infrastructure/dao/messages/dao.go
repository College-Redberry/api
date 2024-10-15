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
        FROM task_management.messages WHERE id=$1
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
	result, err := dao.db.ExecContext(context.Background(), `
        INSERT INTO task_management.messages (
            card_id,
            NOW(),
            user_id,
            parent_message_id
        ) VALUES ($1, $2, $3, $4)
        RETURNING id
    `,
		message.CardID,
		message.UserID,
		message.ParentMessageID,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (dao *Dao) Update(messageID int, updates entities.Message) (int64, error) {
	result, err := dao.db.ExecContext(context.Background(), `
        UPDATE task_management.messages
        SET 
            card_id = COALESCE($2, card_id),
            updated_at = NOW(),
            user_id = COALESCE($3, user_id),
            parent_message_id = COALESCE($4, parent_message_id)
        WHERE id = $1
    `,
		messageID,
		updates.CardID,
		updates.UserID,
		updates.ParentMessageID,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (dao *Dao) Delete(messageID int) (int64, error) {
	result, err := dao.db.ExecContext(context.Background(), `
        DELETE FROM task_management.messages WHERE id = $1
    `, messageID)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
