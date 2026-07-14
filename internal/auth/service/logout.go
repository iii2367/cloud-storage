package service

import (
	"context"
)

func (s *Service) Logout(ctx context.Context, refreshToken string) error {
	claims, err := s.jwtManager.ParseRefreshToken(refreshToken)
	if err != nil {
		return err
	}

	session, err := s.sessionRepo.FindByID(ctx, claims.SessionID)
	if err != nil {
		return err
	}

	if HashToken(refreshToken) != session.TokenHash {
    	return ErrInvalidRefreshToken
	}

	if session.RevokedAt != nil {
		return nil
	}

	if err := s.sessionRepo.Revoke(ctx, session.SessionID); err != nil {
		return err
	}

	return nil
}
