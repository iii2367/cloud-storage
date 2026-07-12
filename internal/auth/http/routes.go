package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *Handler) {
	auth := r.Group("/auth")
	{
		auth.POST("/signup", h.Signup)
		auth.POST("/login", h.Login)
	}

	r.GET("/", func(c *gin.Context) {
		c.File("./web/test_auth.html")
	})
}
