package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie(refreshCookieName)
	if err == nil {
   		_ = h.service.Logout(c.Request.Context(), refreshToken) 
    }

	clearRefreshCookie(c.Writer, c.Request)

	c.Status(http.StatusNoContent)
}
