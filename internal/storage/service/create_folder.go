package service

import (
	"cloud-storage/internal/storage/dto"
	"context"

	"github.com/google/uuid"
)

func (s *Service) CreateFolder(
	ctx context.Context,
	name string,
	description string,
	parentID *uuid.UUID,
	userID uint,
) (*dto.NodeResponse, error) {

	node, err := s.createNode(
		ctx,
		name,
		description,
		"folder",
		parentID,
		userID,
	)
	if err != nil {
		return nil, err
	}

	return toNodeResponse(node), nil
}
