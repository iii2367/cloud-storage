package repository

import (
	"cloud-storage/internal/storage/entity"
	"context"

	"github.com/google/uuid"
)

type TreeNodeRepository interface {
	Create(ctx context.Context, node *entity.TreeNode) error

	FindByID(ctx context.Context, id uuid.UUID, userID uint) (*entity.TreeNode, error)
	FindRoot(ctx context.Context, userID uint) (*entity.TreeNode, error)
	FindChildren(ctx context.Context, parentID *uuid.UUID, userID uint) ([]*entity.TreeNode, error)
	FindByUserID(ctx context.Context, userID uint) ([]*entity.TreeNode, error)

	UpdateName(ctx context.Context, id uuid.UUID, userID uint, name string) error
	UpdateDescription(ctx context.Context, id uuid.UUID, userID uint, description string) error
	UpdateSize(ctx context.Context, id uuid.UUID, userID uint, size int64) error
	UpdateFileMetadata(ctx context.Context, id uuid.UUID, userID uint, extension *string, mimeType *string, size int64) error

	Delete(ctx context.Context, id uuid.UUID, userID uint) error
	DeleteByUserID(ctx context.Context, userID uint) (int64, error)
	DeleteByParentID(ctx context.Context, parentID uuid.UUID, userID uint) (int64, error)
}
