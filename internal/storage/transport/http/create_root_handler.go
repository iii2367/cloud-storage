package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRoot(c *gin.Context) {

	userID := c.MustGet("userID").(uint)


	err := h.service.CreateRoot(
		c.Request.Context(),
		userID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}


	c.JSON(http.StatusCreated, gin.H{
		"message": "root created",
	})
}
