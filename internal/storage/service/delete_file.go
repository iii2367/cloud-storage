package service

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) DeleteFile(
	ctx context.Context,
	nodeID uuid.UUID,
	userID uint,
) (error) {

	node, err := s.treeNodeRepo.FindByID(
		ctx,
		nodeID,
		userID,
	)
	if err != nil {
		return err
	} 

	err = s.deletePhysicalFile(
		node.ID,
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
