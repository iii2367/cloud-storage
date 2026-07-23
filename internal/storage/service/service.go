package service

import (
	"cloud-storage/internal/jwt"
	"cloud-storage/internal/storage/repository"
)

type Service struct {
	treeNodeRepo repository.TreeNodeRepository
	jwtManager   *jwt.Manager
}

func NewService(
	treeNodeRepo repository.TreeNodeRepository,
	jwtManager *jwt.Manager,
) *Service {
	return &Service{
		treeNodeRepo: treeNodeRepo,
		jwtManager:   jwtManager,
	}
}
