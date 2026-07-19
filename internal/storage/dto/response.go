package dto

import (
	"time"

	"github.com/google/uuid"
)

type TreeNodeResponse struct {
	ID          uuid.UUID             `json:"id"`
	ParentID 	*uuid.UUID 			`json:"parent_id,omitempty"`

	Name        string             `json:"name"`
	FileType    string             `json:"file_type"`
	MimeType    string             `json:"mime_type,omitempty"`
	Description string             `json:"description,omitempty"`
	Size        int64              `json:"size"`
	UploadAt    time.Time          `json:"upload_at"`
	UpdatedAt   time.Time          `json:"updated_at"`

	Children []*TreeNodeResponse 	`json:"children,omitempty"`
}

type StorageResponse struct {
	Nodes []*TreeNodeResponse `json:"nodes"`
}
