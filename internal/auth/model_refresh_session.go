package auth

import (
	"time"
	"github.com/google/uuid"
)

type RefreshSession struct {
    SessionID 	uuid.UUID
    UserID 		uint
    TokenHash 	string
    ExpiresAt 	time.Time
    CreatedAt 	time.Time
    LastUsedAt 	time.Time
    Revoked 	bool
    UserAgent 	string
    IPAddress 	string
}
