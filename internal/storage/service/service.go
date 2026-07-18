package service

import (
	"cloud-storage/internal/storage/repository"
	"cloud-storage/internal/jwt"
)

type Service struct {
	treeNodeRepo	repository.TreeNodeRepository
	jwtManager  *jwt.Manager
}

func NewService(
	treeNodeRepo 	repository.TreeNodeRepository,
	jwtManager 		*jwt.Manager,
) *Service {
	return &Service {
		treeNodeRepo: 	treeNodeRepo,
		jwtManager: 	jwtManager,
	}
}
