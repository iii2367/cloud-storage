package web

import ("github.com/gin-gonic/gin")

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	router.GET("/", handler.Index)
	router.GET("/login", handler.Login)
}
