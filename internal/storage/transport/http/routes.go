package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *Handler, m *Middleware) {
	storage := r.Group("/api/storage")
	storage.Use(m.AuthMiddleware())
	{
		storage.GET("/tree", h.GetRootTree) 
		storage.GET("/tree/:id", h.GetTree) 
		storage.POST("/folders", h.CreateFolder)
		storage.POST("/root", h.CreateRoot)    
		storage.POST("/files", h.UploadFile) 
		storage.DELETE("/nodes/:id", h.DeleteNode) 
		storage.GET("/files/:id/download", h.DownloadFile) 
	}
}
