package web

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	router.GET("/", handler.Login)
	router.GET("/signup", handler.Signup)
	router.GET("/storage", handler.Storage)
}
