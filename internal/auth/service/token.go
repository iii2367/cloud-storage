package service

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"time"
)

type TokenPair struct {
	AccessToken      string
	RefreshToken     string
	RefreshExpiresAt time.Time
}

func (s *Service) issueTokens(userID uint, sessionID uuid.UUID) (*TokenPair, error) {
	accessToken, err := s.jwtManager.GenerateAccessToken(userID, sessionID)
	if err != nil {
		return nil, err
	}
	refreshToken, refreshClaims, err := s.jwtManager.GenerateRefreshToken(userID, sessionID)
	if err != nil {
		return nil, err
	}
	return &TokenPair{
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		RefreshExpiresAt: refreshClaims.ExpiresAt.Time,
	}, nil
}

func HashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}
