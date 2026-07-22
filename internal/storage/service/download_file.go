package service

import (
	"context"
	"path/filepath"
	
	"github.com/google/uuid"
)

func (s *Service) DownloadFile(
	ctx context.Context,
	nodeID uuid.UUID,
	userID uint,
) (
	string,
	string,
	string,
	error,
) {

	node, err := s.treeNodeRepo.FindByID(
		ctx,
		nodeID,
		userID,
	)

	if err != nil {
		return "", "", "", err
	}

	if node.FileType != "file" {
		return "",
			"",
			"",
			ErrNotFile	
	}

	path := filepath.Join(
		"storage/files",
		node.ID.String(),
	)
	fileName := node.Name
	if node.Extension != nil {
		fileName += *node.Extension
	}
	mimeType := "application/octet-stream"

	if node.MimeType != nil {
		mimeType = *node.MimeType
	}
	return path, fileName, mimeType, nil
}
