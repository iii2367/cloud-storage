package repository

import (
	"cloud-storage/internal/auth/entity"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByID(ctx context.Context, id uint) (*entity.User, error)
	DeleteByID(ctx context.Context, id uint) error
}
