package service

import (
	"cloud-storage/internal/storage/dto"
	"context"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
)

func (s *Service) UploadFile(
	ctx context.Context,
	file *multipart.FileHeader,
	name string,
	description string,
	parentID *uuid.UUID,
	userID uint,
) (*dto.NodeResponse, error) {

	node, err := s.createNode(
		ctx,
		name,
		description,
		"file",
		parentID,
		userID,
	)
	if err != nil {
		return nil, err
	}

	err = s.saveFile(
		file,
		node.ID,
	)
	if err != nil {
		return nil,err
	}

	extension := filepath.Ext(
		file.Filename,
	)
	mimeType := file.Header.Get(
		"Content-Type",
	)
	err = s.treeNodeRepo.UpdateFileMetadata(
		ctx,
		node.ID,
		userID,
		&extension,
		&mimeType,
		file.Size,
	)
	if err != nil {
		return nil,err
	}
	node.Extension = &extension
	node.MimeType = &mimeType
	node.Size = file.Size

	return toNodeResponse(node), nil
}
