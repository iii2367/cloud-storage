package http

import (	
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {	
		auth := c.GetHeader("Authorization")
		
		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{ "error": "missing token" })
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		
		claims, err := h.jwtManager.ParseAccessToken(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{ "error": "invalid token" })
			return
		}
		c.Set("userID", claims.UserID)

		c.Next()
	}
}
