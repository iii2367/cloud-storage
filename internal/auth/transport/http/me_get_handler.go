package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetMe(c *gin.Context) {

	userID 		:= c.MustGet("userID").(uint)
	sessionID 	:= c.MustGet("sessionID").(uuid.UUID)

	user, err := h.service.GetMe(
		c.Request.Context(),
		userID,
		sessionID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
