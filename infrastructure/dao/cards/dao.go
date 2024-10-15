package dao

import (
	"context"
	"database/sql"
	"time"

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

func (dao *Dao) GetByID(cardID int) (entities.Card, error) {
	var card models.Card

	err := sqlscan.Get(context.Background(), dao.db, &card, `
		SELECT 
			id
			title
			description
			created_at
			start_at
			updated_at
			finished_at
			estimated_finished_at
			status_id
			manager_id
			assigned_id
			priority_id
			parent_card_id
		FROM task_management.cards WHERE id=$1
	`, cardID)
	if err != nil {
		return entities.Card{}, err
	}

	return entities.Card{
		ID:                  card.ID,
		Title:               card.Title,
		Description:         card.Description,
		CreatedAt:           card.CreatedAt,
		StartAt:             card.StartAt.Time,
		UpdatedAt:           card.UpdatedAt.Time,
		FinishedAt:          card.FinishedAt.Time,
		EstimatedFinishedAt: card.EstimatedFinishedAt.Time,
		StatusID:            card.StatusID,
		ManagerID:           card.ManagerID,
		AssignedID:          card.AssignedID,
		PriorityID:          card.PriorityID,
		ParentCardID:        int16(card.ParentCardID.Int64),
	}, nil
}

func (dao *Dao) Create(card entities.Card) (int64, error) {
	result, err := dao.db.ExecContext(context.Background(), `
        INSERT INTO task_management.cards (
            title,
            description,
            NOW(),
            start_at,
            NOW(),
            estimated_finished_at,
            status_id,
            manager_id,
            assigned_id,
            priority_id,
            parent_card_id
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
        RETURNING id
    `,
		card.Title,
		card.Description,
		card.StartAt,
		card.EstimatedFinishedAt,
		card.StatusID,
		card.ManagerID,
		card.AssignedID,
		card.PriorityID,
		card.ParentCardID,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (dao *Dao) Update(cardID int, updates entities.Card) (int64, error) {
	result, err := dao.db.ExecContext(context.Background(), `
        UPDATE task_management.cards
        SET 
            title = COALESCE($2, title),
            description = COALESCE($3, description),
            start_at = COALESCE($4, start_at),
            updated_at = $5,
            estimated_finished_at = COALESCE($6, estimated_finished_at),
            status_id = COALESCE($7, status_id),
            manager_id = COALESCE($8, manager_id),
            assigned_id = COALESCE($9, assigned_id),
            priority_id = COALESCE($10, priority_id),
            parent_card_id = COALESCE($11, parent_card_id)
        WHERE id = $1
    `,
		cardID,
		updates.Title,
		updates.Description,
		updates.StartAt,
		time.Now(),
		updates.EstimatedFinishedAt,
		updates.StatusID,
		updates.ManagerID,
		updates.AssignedID,
		updates.PriorityID,
		updates.ParentCardID,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (dao *Dao) Delete(cardID int) (int64, error) {
	result, err := dao.db.ExecContext(context.Background(), `
        DELETE FROM task_management.cards WHERE id = $1
    `)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
