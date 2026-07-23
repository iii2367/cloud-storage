package service

import (
	"cloud-storage/internal/storage/dto"
	"context"
)

func (s *Service) GetRootTree(
	ctx context.Context,
	userID uint,
) (*dto.TreeNodeResponse, error) {

	root, err := s.treeNodeRepo.FindRoot(ctx, userID)

	if err != nil {
		return nil, err
	}

	children, err := s.treeNodeRepo.FindChildren(ctx, &root.ID, userID)

	if err != nil {
		return nil, err
	}

	return &dto.TreeNodeResponse{
		Node:     *toNodeResponse(root),
		Children: toNodeResponseList(children),
	}, nil
}
