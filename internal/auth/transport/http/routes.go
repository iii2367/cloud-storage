package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *Handler, m *Middleware) {
	auth := r.Group("api/auth")
	{
		auth.POST("/signup", h.Signup)
		auth.POST("/login", h.Login)
		auth.POST("/refresh", h.Refresh)
		auth.POST("/logout", h.Logout)
	}
	users := r.Group("api/users")
	users.Use(m.AuthMiddleware())
	{
		users.GET("/me", h.GetMe)	
		users.DELETE("/me", h.DeleteMe)
	}
}
