package http

import (
	"cloud-storage/internal/jwt"
)

type Middleware struct {
	jwtManager	*jwt.Manager
}

func NewMiddleware(jwtManager *jwt.Manager) *Middleware {
	return &Middleware {
		jwtManager: jwtManager,	
	}
}
