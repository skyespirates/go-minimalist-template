package repository

import (
	"context"

	"github.com/skyespirates/go-minimalist-template/internal/entity"
)

type TaskRepository interface {
	GetAll(context.Context) ([]*entity.Task, error)
	GetById(context.Context, int) (*entity.Task, error)
}
