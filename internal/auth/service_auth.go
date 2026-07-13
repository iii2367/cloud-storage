package auth

import (
	"cloud-storage/internal/auth/dto"
	"cloud-storage/internal/jwt"
	"context"
	"time"
	"net"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userRepo	UserRepository
	sessionRepo SessionRepository
	jwtManager  *jwt.Manager
}

type LoginMeta struct {
    UserAgent string
    IP        net.IP
}

func NewService(
	userRepo 	UserRepository,
	sessionRepo SessionRepository,
	jwtManager 	*jwt.Manager,
) *Service {
	return &Service {
		userRepo: 		userRepo,
		sessionRepo: 	sessionRepo,
		jwtManager: 	jwtManager,
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

func (s *Service) Login(ctx context.Context, email, password string, meta LoginMeta) (*dto.LoginResponse, *TokenPair, error) {

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
	sid := uuid.New()
	tokens, err := s.issueTokens(user.ID, sid)
	if err != nil {
		return nil, nil, err
	}

	session := &Session{
		SessionID: 	sid,
		UserID: 	user.ID,
		TokenHash: 	HashToken(tokens.RefreshToken),
		ExpiresAt: 	tokens.RefreshClaims.ExpiresAt.Time,
		UserAgent: 	meta.UserAgent,
		IPAddress:  meta.IP,
	}
	err = s.sessionRepo.Create(ctx, session)
	if err != nil {
		return nil, nil, err
	}
	
	return &dto.LoginResponse{AccessToken: tokens.AccessToken}, tokens, nil
}
