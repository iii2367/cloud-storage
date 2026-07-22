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

	// видалення фізичного файлу 


	err := s.treeNodeRepo.Delete(ctx, nodeID, userID)
	if err != nil {
		return err
	}

	return nil
}
