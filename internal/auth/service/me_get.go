package service

import (
	"cloud-storage/internal/auth/dto"
	"context"

	"github.com/google/uuid"
)


func (s *Service) GetMe(ctx context.Context, userID uint, sessionID uuid.UUID) (*dto.UserResponse, error) {	
	err := s.sessionRepo.UpdateLastUsedAt(ctx, sessionID)
	if err != nil {
		return  nil, err
	}

	user, err := s.userRepo.FindByID(ctx,userID)

	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		Name: user.Name,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}
