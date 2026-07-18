package http

import (	
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {	
		auth := c.GetHeader("Authorization")
		
		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ "error": "missing token" })
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		
		claims, err := m.jwtManager.ParseAccessToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ "error": "invalid token" })
			return
		}
		c.Set("userID", claims.UserID)
		c.Set("sessionID", claims.SessionID)

		c.Next()
	}
}
