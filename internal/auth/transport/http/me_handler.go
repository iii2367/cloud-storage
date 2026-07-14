package http

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetMe(c *gin.Context) {

	userID := c.MustGet("userID").(uint)

	user, err := h.service.GetMe(
		c.Request.Context(),
		userID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
