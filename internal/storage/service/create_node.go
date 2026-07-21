package service

import (
	"cloud-storage/internal/storage/dto"
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

func nodeToResponse(node *entity.TreeNode) *dto.NodeResponse {
	return &dto.NodeResponse{
		ID:          node.ID,
		ParentID:    node.ParentID,
		Name:        node.Name,
		FileType:    node.FileType,
		Extension:   node.Extension,
		MimeType:    node.MimeType,
		Description: node.Description,
		Size:        node.Size,
		UploadAt:    node.UploadAt,
		UpdatedAt:   node.UpdatedAt,
	}
}
