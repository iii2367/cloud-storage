package service

import (
	"cloud-storage/internal/storage/dto"
	"cloud-storage/internal/storage/entity"
	"context"
	"time"
	
	"github.com/google/uuid"
)


func (s *Service) CreateNode(
	ctx 		context.Context, 
	name 		string,
	description string,
	fileType	string,
	parentID 	*uuid.UUID,
	userID 		uint,
) (*dto.NodeResponse, error) {	

	node := &entity.TreeNode{
		ID:				uuid.New(),
		ParentID: 		parentID,
		UserID:		 	userID,
		FileType: 		fileType,
		Extension: 		nil,
		MimeType:	 	nil,
		Name:  			name,
		Description: 	description,
		Size:			0,
	}	

	err := s.treeNodeRepo.Create(ctx, node)
	if err != nil {
		return nil, err
	}

	return &dto.NodeResponse{
		ID: 			node.ID,
		ParentID: 		node.ParentID,
		Name:			node.Name,
		FileType:		node.FileType,
		Extension: 		node.Extension,
		MimeType:		node.MimeType,
		Description: 	node.Description,
		Size:			node.Size,
		UploadAt: 		time.Now(),
		UpdatedAt:		time.Now(),
	}, nil
}
