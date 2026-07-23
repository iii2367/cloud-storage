package service

import (
	"cloud-storage/internal/storage/dto"
	"context"

	"github.com/google/uuid"
)

func (s *Service) GetTree(
	ctx context.Context,
	treeID uuid.UUID,
	userID uint,
) (*dto.TreeNodeResponse, error) {

	node, err := s.treeNodeRepo.FindByID(ctx, treeID, userID)

	if err != nil {
		return nil, err
	}

	children, err := s.treeNodeRepo.FindChildren(ctx, &treeID, userID)

	if err != nil {
		return nil, err
	}

	return &dto.TreeNodeResponse{
		Node:     *toNodeResponse(node),
		Children: toNodeResponseList(children),
	}, nil
}
