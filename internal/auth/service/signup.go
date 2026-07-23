package service

import (
	"cloud-storage/internal/auth/dto"
	"cloud-storage/internal/auth/entity"
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Signup(ctx context.Context, name, email, password string) (*dto.UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &entity.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(hash),
	}
	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return &dto.UserResponse{Name: name, Email: email, CreatedAt: time.Now()}, nil
}
