package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteMe(c *gin.Context) {

	userID := c.MustGet("userID").(uint)

	err := h.service.DeleteMe(
		c.Request.Context(),
		userID,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	clearRefreshCookie(
		c.Writer,
		c.Request,
	)

	c.Status(http.StatusNoContent)
}
