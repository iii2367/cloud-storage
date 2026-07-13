package service

import (
	"cloud-storage/internal/auth/repository"
	"cloud-storage/internal/jwt"
	"net"
)

type Service struct {
	userRepo	repository.UserRepository
	sessionRepo repository.SessionRepository
	jwtManager  *jwt.Manager
}

type LoginMeta struct {
    UserAgent string
    IP        net.IP
}

func NewService(
	userRepo 	repository.UserRepository,
	sessionRepo repository.SessionRepository,
	jwtManager 	*jwt.Manager,
) *Service {
	return &Service {
		userRepo: 		userRepo,
		sessionRepo: 	sessionRepo,
		jwtManager: 	jwtManager,
	}
}
