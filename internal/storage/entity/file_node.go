package entity

import (
	"time"

	"github.com/google/uuid"
)

type TreeNode struct {
	ID			uuid.UUID	
	ParentID	*uuid.UUID 	
	UserID		uint		

	FileType	string		
	MimeType 	string

	UploadAt	time.Time	
	UpdatedAt 	time.Time	

	Name		string		
	Description	string		

	Size 		int64 		
}
