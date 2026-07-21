package dto

import ("github.com/google/uuid")

type CreateFolderRequest struct {
	Name        string     `json:"name" binding:"required"`
	Description string     `json:"description"`
	ParentID    *uuid.UUID `json:"parent_id"`
}

type UpdateNodeRequest struct {
	Name 		*string `json:"name"`
	Description *string `json:"description"`
}
