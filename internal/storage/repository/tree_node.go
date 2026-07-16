package repository

import (
	"cloud-storage/internal/storage/entity"
	"context"

	"github.com/google/uuid"
)

type TreeNodeRepository interface {
	Create(ctx context.Context, node *entity.TreeNode) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.TreeNode, error)
	FindByUserID(ctx context.Context, id uint) ([]*entity.TreeNode, error)
	FindByParentID(ctx context.Context, id uuid.UUID) ([]*entity.TreeNode, error)
	UpdateName(ctx context.Context, id uuid.UUID, name string) error
	UpdateDescription(ctx context.Context, id uuid.UUID, description string) error
	UpdateSize(ctx context.Context, id uuid.UUID, size int64) error
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteByUserID(ctx context.Context, id uint) (int64, error)
	DeleteByParentID(ctx context.Context, id uuid.UUID) (int64, error)	
}
