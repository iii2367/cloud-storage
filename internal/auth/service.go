package auth

import (
	"time"
	"context"
	"cloud-storage/internal/auth/dto"
	"cloud-storage/internal/jwt"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

type Service struct {
	userRepo	UserRepository
	sessionRepo SessionRepository
	jwtManager  *jwt.Manager
}

func NewService(
	userRepo 	UserRepository,
	sessionRepo SessionRepository,
	jwtManager 	*jwt.Manager,
) *Service {
	return &Service {
		userRepo: 	userRepo,
		jwtManager: jwtManager,
	}
}

func (s *Service) Signup(ctx context.Context, name, email, password string) (*dto.SignupResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &User {
		Name:         name,
		Email:        email,
		PasswordHash: string(hash),
	}
	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return &dto.SignupResponse{Name: name, Email: email, CreatedAt: time.Now()}, nil
} 

func (s *Service) Login(ctx context.Context, email, password string) (*dto.LoginResponse, *tokenPair, error) {

	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, nil, err
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)
	if err != nil {
		return nil, nil, ErrInvalidCredentials
	}
	tokens, err := s.issueTokens(ctx, user.ID, uuid.New())
	
	return &dto.LoginResponse{AccessToken: tokens.AccessToken}, tokens, nil
}

type tokenPair struct {
    AccessToken  string
    RefreshToken string
}

func (s *Service) issueTokens(ctx context.Context, userID uint, sessionID uuid.UUID) (*tokenPair, error) {
	accessToken, err := s.jwtManager.GenerateAccessToken(userID)
	if err != nil {
		return  nil, err
	}
	refreshToken, err := s.jwtManager.GenerateRefreshToken(userID, sessionID)
	if err != nil {
		return nil, err
	}
	return &tokenPair{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
