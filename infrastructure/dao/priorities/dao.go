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

func (dao *Dao) GetByID(priorityID int) (entities.Priority, error) {
	var priority models.Priority

	err := sqlscan.Get(context.Background(), dao.db, &priority, `
        SELECT 
            id,
            name,
            color
        FROM task_management.priorities WHERE id=$1
    `, priorityID)
	if err != nil {
		return entities.Priority{}, err
	}

	return entities.Priority{
		ID:    priority.ID,
		Name:  priority.Name,
		Color: priority.Color,
	}, nil
}

func (dao *Dao) Create(priority entities.Priority) (int64, error) {
	result, err := dao.db.ExecContext(context.Background(), `
        INSERT INTO task_management.priorities (
            name,
            color
        ) VALUES ($1, $2)
        RETURNING id
    `,
		priority.Name,
		priority.Color,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (dao *Dao) Update(priorityID int, updates entities.Priority) (int64, error) {
	result, err := dao.db.ExecContext(context.Background(), `
        UPDATE task_management.priorities
        SET 
            name = COALESCE($2, name),
            color = COALESCE($3, color)
        WHERE id = $1
    `,
		priorityID,
		updates.Name,
		updates.Color,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (dao *Dao) Delete(priorityID int) (int64, error) {
	result, err := dao.db.ExecContext(context.Background(), `
        DELETE FROM task_management.priorities WHERE id = $1
    `, priorityID)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
