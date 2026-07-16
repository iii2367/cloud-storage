package dto

import "time"

type TreeNodeResponse struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	FileType    string             `json:"file_type"`
	MimeType    string             `json:"mime_type,omitempty"`
	Description string             `json:"description,omitempty"`
	Size        int64              `json:"size"`
	UploadAt    time.Time          `json:"upload_at"`
	UpdatedAt   time.Time          `json:"updated_at"`

	Children []*TreeNodeResponse `json:"children,omitempty"`
}

type StorageResponse struct {
	Nodes []*TreeNodeResponse `json:"nodes"`
}
