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

func (dao *Dao) GetByID(statusID int) (entities.Status, error) {
	var status models.Status

	err := sqlscan.Get(context.Background(), dao.db, &status, `
        SELECT 
            id,
            name,
            color
        FROM task_management.statuses WHERE id=$1
    `, statusID)
	if err != nil {
		return entities.Status{}, err
	}

	return entities.Status{
		ID:    status.ID,
		Name:  status.Name,
		Color: status.Color,
	}, nil
}

func (dao *Dao) Create(status entities.Status) (int64, error) {
	var id int64

	err := sqlscan.Get(context.Background(), dao.db, &id, `        
		INSERT INTO task_management.statuses (
            name,
            color
        ) VALUES ($1, $2)
        RETURNING id
    `,
		status.Name,
		status.Color,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *Dao) Update(statusID int, updates entities.Status) (int64, error) {
	var id int64

	err := sqlscan.Get(context.Background(), dao.db, &id, `        
		UPDATE task_management.statuses
        SET 
            name = COALESCE($2, name),
            color = COALESCE($3, color)
        WHERE id = $1 RETURNING id
    `,
		statusID,
		updates.Name,
		updates.Color,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *Dao) Delete(statusID int) (int64, error) {
	var id int64

	err := sqlscan.Get(context.Background(), dao.db, &id, `        
		DELETE FROM task_management.statuses WHERE id = $1 RETURNING id
    `, statusID)
	if err != nil {
		return 0, err
	}

	return id, nil
}
