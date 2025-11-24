package pgsql

import (
	"context"
	"database/sql"
	"time"

	"github.com/skyespirates/go-minimalist-template/internal/entity"
	"github.com/skyespirates/go-minimalist-template/internal/repository"
)

type taskRepo struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) repository.TaskRepository {
	return &taskRepo{
		db: db,
	}
}

func (tp *taskRepo) GetAll(ctx context.Context) ([]*entity.Task, error) {
	query := `SELECT * FROM tasks`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := tp.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*entity.Task

	for rows.Next() {
		var task entity.Task
		err := rows.Scan(&task.Id, &task.Title, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tp *taskRepo) GetById(ctx context.Context, id int) (*entity.Task, error) {
	query := `SELECT * FROM tasks WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var task entity.Task

	err := tp.db.QueryRowContext(ctx, query, id).Scan(&task.Id, &task.Title, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
