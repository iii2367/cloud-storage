package service

import (
	"cloud-storage/internal/storage/entity"
	"context"

	"github.com/google/uuid"
)

func (s *Service) createNode(
	ctx context.Context,
	name string,
	description string,
	fileType string,
	parentID *uuid.UUID,
	userID uint,
) (*entity.TreeNode, error) {

	node := &entity.TreeNode{
		ID:          uuid.New(),
		ParentID:    parentID,
		UserID:      userID,
		FileType:    fileType,
		Name:        name,
		Description: description,
		Size:        0,
	}

	if err := s.treeNodeRepo.Create(ctx, node); err != nil {
		return nil, err
	}

	return node, nil
}
