package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie(refreshCookieName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "refresh token not found",
		})
		return
	}
	tokenResponse, tokens, err := h.service.Refresh(c.Request.Context(), refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}
	setRefreshCookie(
		c.Writer,
		c.Request,
		tokens.RefreshToken,
		tokens.RefreshExpiresAt,
	)
	c.JSON(http.StatusOK, tokenResponse)
}
