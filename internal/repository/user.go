package repository

import (
	"context"

	"github.com/skyespirates/go-minimalist-template/internal/entity"
)

type UserRepository interface {
	Create(context.Context, entity.RegisterPayload) (*entity.User, error)
	FindByEmail(context.Context, string) (*entity.User, error)
}
