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

func (dao *Dao) GetByID(cardID int) (entities.Card, error) {
	var card models.Card

	err := sqlscan.Get(context.Background(), dao.db, &card, `
		SELECT 
			id,
			title,
			description,
			created_at,
			start_at,
			updated_at,
			finished_at,
			estimated_finished_at,
			status_id,
			manager_id,
			assigned_id,
			priority_id,
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
	var id int64

	var parentCardID *int16
	if card.ParentCardID != 0 {
		parentCardID = &card.ParentCardID
	}

	err := sqlscan.Get(context.Background(), dao.db, &id, `        
		INSERT INTO task_management.cards (
            title,
            description,
            start_at,
            estimated_finished_at,
            status_id,
            manager_id,
            assigned_id,
            priority_id,
            parent_card_id
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
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
		parentCardID,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *Dao) Update(cardID int, updates entities.Card) (int64, error) {
	var id int64

	err := sqlscan.Get(context.Background(), dao.db, &id, `        
		UPDATE task_management.cards
        SET 
            title = COALESCE($2, title),
            description = COALESCE($3, description),
            start_at = COALESCE($4, start_at),
            estimated_finished_at = COALESCE($5, estimated_finished_at),
            status_id = COALESCE($6, status_id),
            manager_id = COALESCE($7, manager_id),
            assigned_id = COALESCE($8, assigned_id),
            priority_id = COALESCE($9, priority_id),
            parent_card_id = COALESCE($10, parent_card_id),
            updated_at = NOW()
        WHERE id = $1 RETURNING id
    `,
		cardID,
		updates.Title,
		updates.Description,
		updates.StartAt,
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

	return id, nil
}

func (dao *Dao) Delete(cardID int) (int64, error) {
	var id int64

	err := sqlscan.Get(context.Background(), dao.db, &id, `        
		DELETE FROM task_management.cards WHERE id = $1 RETURNING id
    `, cardID)
	if err != nil {
		return 0, err
	}

	return id, nil
}
