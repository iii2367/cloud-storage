package service

import (	
	"cloud-storage/internal/auth/dto"
	"context"
)


func (s *Service) GetMe(ctx context.Context, userID uint) (*dto.UserResponse, error) {	
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
