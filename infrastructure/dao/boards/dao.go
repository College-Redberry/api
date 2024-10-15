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

func (dao *Dao) GetByID(boardID int) (entities.Board, error) {
	var board models.Board

	err := sqlscan.Get(context.Background(), dao.db, &board, `
        SELECT 
            id,
            name,
            description,
            manager_id,
            created_at,
            updated_at,
            project_id
        FROM task_management.boards WHERE id=$1
    `, boardID)
	if err != nil {
		return entities.Board{}, err
	}

	return entities.Board{
		ID:          board.ID,
		Name:        board.Name,
		Description: board.Description,
		ManagerID:   board.ManagerID,
		CreatedAt:   board.CreatedAt,
		UpdatedAt:   board.UpdatedAt.Time,
		ProjectID:   board.ProjectID,
	}, nil
}

func (dao *Dao) Create(board entities.Board) (int64, error) {
	result, err := dao.db.ExecContext(context.Background(), `
        INSERT INTO task_management.boards (
            name,
            description,
            manager_id,
            NOW()
        ) VALUES ($1, $2, $3, $4)
        RETURNING id
    `,
		board.Name,
		board.Description,
		board.ManagerID,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (dao *Dao) Update(boardID int, updates entities.Board) (int64, error) {
	result, err := dao.db.ExecContext(context.Background(), `
        UPDATE task_management.boards
        SET 
            name = COALESCE($2, name),
            description = COALESCE($3, description),
            manager_id = COALESCE($4, manager_id),
            updated_at = $5
        WHERE id = $1
    `,
		boardID,
		updates.Name,
		updates.Description,
		updates.ManagerID,
		time.Now(),
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (dao *Dao) Delete(boardID int) (int64, error) {
	result, err := dao.db.ExecContext(context.Background(), `
        DELETE FROM task_management.boards WHERE id = $1
    `)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
