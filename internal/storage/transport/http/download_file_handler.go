package http

import (		
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) DownloadFile(c *gin.Context) {

	userID := c.MustGet("userID").(uint)


	nodeID, err := uuid.Parse(
		c.Param("id"),
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid file id",
		})
		return
	}

	filePath, fileName, mimeType, err :=
		h.service.DownloadFile(
			c.Request.Context(),
			nodeID,
			userID,
		)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.Header(
		"Content-Disposition",
		"attachment; filename=\""+fileName+"\"",
	)

	c.Header(
		"Content-Type",
		mimeType,
	)

	c.File(filePath)
}
