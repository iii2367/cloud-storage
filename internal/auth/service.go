package auth

import (
	"context"
	"cloud-storage/internal/auth/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo	Repository
	jwt  	jwt.Manager
}

func NewService(repo Repository, jwt jwt.Manager) *Service {
	return &Service {
		repo: 	repo,
		jwt: 	jwt,
	}
}

func (s *Service) Register(ctx context.Context, name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &User {
		Name:         name,
		Email:        email,
		PasswordHash: string(hash),
	}
	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
} 

func (s *Service) Login(ctx context.Context, email, password string) (*User, error) {

	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)
	if err != nil {
		return nil, ErrInvalidCredentials
	}
	return user, nil
}
