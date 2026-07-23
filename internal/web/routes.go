package web

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	router.LoadHTMLGlob("web/templates/*")

	router.Static("/css", "./web/static/css")
	router.Static("/js/core", "./web/static/js/core")
	router.Static("/js/storage", "./web/static/js/storage")

	router.GET("/", handler.Login)
	router.GET("/signup", handler.Signup)
	router.GET("/storage", handler.Storage)
}
