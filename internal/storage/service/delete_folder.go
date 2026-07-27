package service

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) DeleteFolder(
	ctx context.Context,
	nodeID uuid.UUID,
	userID uint,
) error {

	children, err := s.treeNodeRepo.FindChildren(ctx, &nodeID, userID)
	if err != nil {
		return err
	}

	for _, child := range children {

		switch child.FileType {

		case "file":
			if err := s.DeleteFile(ctx, child.ID, userID); err != nil {
				return err
			}

		case "folder":
			if err := s.DeleteFolder(ctx, child.ID, userID); err != nil {
				return err
			}
		}
	}

	return s.treeNodeRepo.Delete(ctx, nodeID, userID)
}
