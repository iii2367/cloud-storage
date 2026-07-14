package service

import (
	"context"
)

func (s *Service) DeleteMe(ctx context.Context, userID uint) error {

	err := s.sessionRepo.RevokeAllByUserID(ctx, userID)

	if err != nil {
		return err
	}

	err = s.userRepo.DeleteByID(ctx, userID)

	if err != nil {
		return err
	}

	return nil
}
