package dto

type CreateNodeRequest struct {
	Name 		string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	FileType 	string  `json:"file_type" binding:"required"`
	ParentID 	*string `json:"parent_id"`
}

type UpdateNodeRequest struct {
	Name 		*string `json:"name"`
	Description *string `json:"description"`
}
