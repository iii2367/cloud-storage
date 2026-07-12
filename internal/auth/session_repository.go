package auth

import (
	"context"
	"github.com/google/uuid"
)

type SessionRepository interface {
	Create(ctx context.Context, session *Session) error
	FindByID(ctx context.Context, id uuid.UUID) (*Session, error)
	FindByTokenHash(ctx context.Context, hash string) (*Session, error)
	FindByUserID(ctx context.Context, id uint) ([]*Session, error)
	UpdateLastUsedAt(ctx context.Context, id uuid.UUID) error
	Revoke(ctx context.Context, id uuid.UUID) error
	RevokeAllByUserID(ctx context.Context, id uint) error
	DeleteExpired(ctx context.Context) error
}
