package service

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) DeleteNode(
	ctx context.Context,
	nodeID uuid.UUID,
	userID uint,
) (error) {

	node, err := s.treeNodeRepo.FindByID(ctx, nodeID, userID)
	if err != nil {
		return err
	}

	if node.FileType == "folder" {
		err = s.DeleteFolder(ctx, nodeID, userID)
		return err
	} else if node.FileType == "file" {
		err = s.DeleteFile(ctx, nodeID, userID)
		return err
	}

	return nil
}
