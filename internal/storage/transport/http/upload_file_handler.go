package http

import (	
	"cloud-storage/internal/storage/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) UploadFile(c *gin.Context) {

	userID := c.MustGet("userID").(uint)

	fileHeader, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "file is required",
		})
		return
	}

	parentIDString := c.PostForm("parent_id")
	var parentID *uuid.UUID

	if parentIDString != "" {
		id, err := uuid.Parse(parentIDString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid parent_id",
			})
			return
		}
		parentID = &id
	}

	req := dto.UploadFileRequest{
		Name: c.PostForm("name"),
		Description: c.PostForm("description"),
		ParentID: parentID,
	}

	if req.Name == "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "name is required",
		})

		return
	}

	node, err := h.service.UploadFile(
		c.Request.Context(),
		fileHeader,
		req.Name,
		req.Description,
		req.ParentID,
		userID,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated,node)
}
