package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AccessClaims struct {
	UserID		uint	`json:"user_id"` 
	SessionID uuid.UUID `json:"sid"`

    jwt.RegisteredClaims
}

type RefreshClaims struct {
	UserID    uint      `json:"user_id"`
    SessionID uuid.UUID `json:"sid"`

    jwt.RegisteredClaims
}
