package repository

import (
	"cloud-storage/internal/auth/entity"
	"context"
	"github.com/google/uuid"
)

type SessionRepository interface {
	Create(ctx context.Context, session *entity.Session) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Session, error)
	FindByTokenHash(ctx context.Context, hash string) (*entity.Session, error)
	FindByUserID(ctx context.Context, id uint) ([]*entity.Session, error)
	UpdateTokenHash(ctx context.Context, id uuid.UUID, hash string) error
	UpdateLastUsedAt(ctx context.Context, id uuid.UUID) error
	Revoke(ctx context.Context, id uuid.UUID) error
	RevokeAllByUserID(ctx context.Context, id uint) error
	DeleteExpired(ctx context.Context) (int64, error)
}
