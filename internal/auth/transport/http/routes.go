package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *Handler, m *Middleware) {
	auth := r.Group("/auth")
	{
		auth.POST("/signup", h.Signup)
		auth.POST("/login", h.Login)
	
		auth.POST("/refresh", h.Refresh)

		auth.GET("/me", m.AuthMiddleware(), h.GetMe)
	}
}
