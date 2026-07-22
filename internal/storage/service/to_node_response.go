package service

import (
	"cloud-storage/internal/storage/dto"
	"cloud-storage/internal/storage/entity"
)

func toNodeResponse(node *entity.TreeNode) *dto.NodeResponse {
	if node == nil {
		return nil
	}

	return &dto.NodeResponse{
		ID:          node.ID,
		ParentID:    node.ParentID,
		Name:        node.Name,
		FileType:    node.FileType,
		Extension:   node.Extension,
		MimeType:    node.MimeType,
		Description: node.Description,
		Size:        node.Size,
		UploadAt:    node.UploadAt,
		UpdatedAt:   node.UpdatedAt,
	}
}

func toNodeResponseList(nodes []*entity.TreeNode) []*dto.NodeResponse {
	result := make([]*dto.NodeResponse, 0, len(nodes))

	for _, node := range nodes {
		result = append(result, toNodeResponse(node))
	}

	return result
}
