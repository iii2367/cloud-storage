package http

import (
	"cloud-storage/internal/storage/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateNode(c *gin.Context) {
	
	userID 		:= c.MustGet("userID").(uint)

	var req dto.CreateNodeRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	node, err := h.service.CreateNode(
		c.Request.Context(),
		req.Name,
		req.Description,
		req.FileType,
		req.ParentID,
		userID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, node)
}
