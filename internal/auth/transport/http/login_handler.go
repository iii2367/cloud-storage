package http

import (
	"cloud-storage/internal/auth/dto"
	"cloud-storage/internal/auth/service"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var req dto.LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokenResponse, tokens, err := h.service.Login(
		c.Request.Context(),
		req.Email,
		req.Password,
		service.LoginMeta{
			UserAgent: c.Request.UserAgent(),
			IP:        net.ParseIP(c.ClientIP()),
		},
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
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
