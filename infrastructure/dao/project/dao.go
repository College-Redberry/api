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

func (dao *Dao) GetByID(projectID int) (entities.Project, error) {
	var project models.Project

	err := sqlscan.Get(context.Background(), dao.db, &project, `
        SELECT 
            id,
            name,
            description,
            created_at,
            updated_at
        FROM task_management.projects WHERE id=$1
    `, projectID)
	if err != nil {
		return entities.Project{}, err
	}

	return entities.Project{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt.Time,
	}, nil
}

func (dao *Dao) Create(project entities.Project) (int64, error) {
	var id int64

	err := sqlscan.Get(context.Background(), dao.db, &id, `        
		INSERT INTO task_management.projects (
            name,
            description
        ) VALUES ($1, $2)
        RETURNING id
    `,
		project.Name,
		project.Description,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *Dao) Update(projectID int, updates entities.Project) (int64, error) {
	var id int64

	err := sqlscan.Get(context.Background(), dao.db, &id, `        
		UPDATE task_management.projects
        SET 
            name = COALESCE($2, name),
            description = COALESCE($3, description),
            updated_at = NOW()
        WHERE id = $1 RETURNING id
    `,
		projectID,
		updates.Name,
		updates.Description,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *Dao) Delete(projectID int) (int64, error) {
	var id int64

	err := sqlscan.Get(context.Background(), dao.db, &id, `        
		DELETE FROM task_management.projects WHERE id = $1 RETURNING id
    `, projectID)
	if err != nil {
		return 0, err
	}

	return id, nil
}
