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
) (*dto.TreeNodeResponse, error) {	

	node := &entity.TreeNode{
		ID:				uuid.New(),
		ParentID: 		parentID,
		UserID:		 	userID,
		FileType: 		fileType,
		MimeType:	 	"empty",
		Name:  			name,
		Description: 	description,
		Size:			0,
	}	

	err := s.treeNodeRepo.Create(ctx, node)
	if err != nil {
		return nil, err
	}

	return &dto.TreeNodeResponse{
		ID: 			node.ID,
		ParentID: 		node.ParentID,
		Name:			node.Name,
		FileType:		node.FileType,
		MimeType:		node.MimeType,
		Description: 	node.Description,
		Size:			node.Size,
		UploadAt: 		time.Now(),
		UpdatedAt:		time.Now(),
		Children: 		nil,
	}, nil
}
