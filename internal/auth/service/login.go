package service

import (	
	"cloud-storage/internal/auth/dto"
	"cloud-storage/internal/auth/entity"
	"context"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

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

	session := &entity.Session{
		SessionID: 	sid,
		UserID: 	user.ID,
		TokenHash: 	HashToken(tokens.RefreshToken),
		ExpiresAt: 	tokens.RefreshExpiresAt,
		UserAgent: 	meta.UserAgent,
		IPAddress:  meta.IP,
	}
	err = s.sessionRepo.Create(ctx, session)
	if err != nil {
		return nil, nil, err
	}
	
	return &dto.LoginResponse{AccessToken: tokens.AccessToken}, tokens, nil
}
