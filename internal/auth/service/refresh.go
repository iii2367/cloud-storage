package service

import (
	"cloud-storage/internal/auth/dto"
	"cloud-storage/internal/auth/entity"
	"context"
	"github.com/google/uuid"
	"time"
)

func (s *Service) Refresh(ctx context.Context, refreshToken string) (*dto.TokenResponse, *TokenPair, error) {
	claims, err := s.jwtManager.ParseRefreshToken(refreshToken)
	if err != nil {
		return nil, nil, err
	}

	session, err := s.sessionRepo.FindByID(ctx, claims.SessionID)
	if err != nil {
		return nil, nil, err
	}

	if HashToken(refreshToken) != session.TokenHash {
		return nil, nil, ErrInvalidRefreshToken
	}

	if session.RevokedAt != nil {
		err = s.sessionRepo.RevokeAllByUserID(ctx, session.UserID)
		if err != nil {
			return nil, nil, err
		}
		return nil, nil, ErrRefreshTokenReuse
	}

	if time.Now().After(session.ExpiresAt) {
		err = s.sessionRepo.Revoke(ctx, session.SessionID)
		if err != nil {
			return nil, nil, err
		}
		return nil, nil, ErrInvalidRefreshToken
	}

	if err := s.sessionRepo.Revoke(ctx, session.SessionID); err != nil {
		return nil, nil, err
	}

	newSessionID := uuid.New()

	tokens, err := s.issueTokens(session.UserID, newSessionID)
	if err != nil {
		return nil, nil, err
	}

	newSession := &entity.Session{
		SessionID: newSessionID,
		UserID:    session.UserID,
		TokenHash: HashToken(tokens.RefreshToken),
		ExpiresAt: tokens.RefreshExpiresAt,
		UserAgent: session.UserAgent,
		IPAddress: session.IPAddress,
	}

	if err := s.sessionRepo.Create(ctx, newSession); err != nil {
		return nil, nil, err
	}

	return &dto.TokenResponse{AccessToken: tokens.AccessToken}, tokens, nil
}
