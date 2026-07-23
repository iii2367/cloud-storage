package service

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) DeleteFile(
	ctx context.Context,
	nodeID uuid.UUID,
	userID uint,
) error {

	err := s.deletePhysicalFile(
		nodeID,
	)
	if err != nil {

		return err
	}

	err = s.treeNodeRepo.Delete(ctx, nodeID, userID)
	if err != nil {
		return err
	}

	return nil
}
