package usecase

import (
	"context"
	"strconv"

	"github.com/skyespirates/go-minimalist-template/internal/entity"
	"github.com/skyespirates/go-minimalist-template/internal/repository"
)

type TaskUsecase interface {
	GetAll(context.Context) ([]*entity.Task, error)
	GetById(context.Context, string) (*entity.Task, error)
}

type taskUsecase struct {
	repo repository.TaskRepository
}

func NewTaskUsecase(repo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{
		repo: repo,
	}
}

func (tu *taskUsecase) GetAll(ctx context.Context) ([]*entity.Task, error) {
	tasks, err := tu.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tu *taskUsecase) GetById(ctx context.Context, id string) (*entity.Task, error) {
	task_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	task, err := tu.repo.GetById(ctx, task_id)
	if err != nil {
		return nil, err
	}
	return task, nil
}
