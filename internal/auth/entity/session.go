package entity

import (
	"time"
	"net"
	"github.com/google/uuid"
)

type Session struct {
    SessionID 	uuid.UUID
    UserID 		uint
    TokenHash 	string
    ExpiresAt 	time.Time
    CreatedAt	time.Time
    LastUsedAt 	time.Time
    RevokedAt 	*time.Time
    UserAgent 	string
    IPAddress 	net.IP
}
