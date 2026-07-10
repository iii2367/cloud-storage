package auth

import (
	"cloud-storage/internal/auth/dto"
	"cloud-storage/internal/jwt"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userRepo	UserRepository
	jwt  	*jwt.Manager
}

func NewService(userRepo UserRepository, jwt *jwt.Manager) *Service {
	return &Service {
		userRepo: 	userRepo,
		jwt: 	jwt,
	}
}

func (s *Service) Signup(ctx context.Context, name, email, password string) (error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &User {
		Name:         name,
		Email:        email,
		PasswordHash: string(hash),
	}
	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
} 

func (s *Service) Login(ctx context.Context, email, password string) (*dto.LoginResponse, error) {

	user, err := s.userRepo.FindByEmail(ctx, email)
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
	accessToken, err := s.jwt.GenerateAccessToken(user.ID)
	if err != nil {
		return  nil, err
	}
	/*refreshToken, refreshClaims, err := s.jwt.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}*/
	return &dto.LoginResponse{AccessToken: accessToken, RefreshToken: "gfff"/*refreshToken*/}, nil
}
