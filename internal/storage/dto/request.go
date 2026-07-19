package dto

import ("github.com/google/uuid")

type CreateNodeRequest struct {
	Name 		string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	FileType 	string  `json:"file_type" binding:"required"`
	ParentID 	*uuid.UUID `json:"parent_id"`
}

type UpdateNodeRequest struct {
	Name 		*string `json:"name"`
	Description *string `json:"description"`
}
