package entity

import (
	"github.com/google/uuid"
	"net"
	"time"
)

type Session struct {
	SessionID  uuid.UUID
	UserID     uint
	TokenHash  string
	ExpiresAt  time.Time
	CreatedAt  time.Time
	LastUsedAt time.Time
	RevokedAt  *time.Time
	UserAgent  string
	IPAddress  net.IP
}
