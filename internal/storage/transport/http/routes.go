package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *Handler, m *Middleware) {
	storage := r.Group("/api/storage")
	storage.Use(m.AuthMiddleware())
	{
/*		// отримання всіх нод
		storage.GET("/tree", h.GetTree)
		// отримати кореневі ноди користувача
		storage.GET("/root", h.GetRoot)
		// отримати одну ноду
		storage.GET("/node/:id", h.GetNode)
		// створити файл або папку
		*/storage.POST("/node", h.CreateNode)/*
		// змінити назву/опис
		storage.PATCH("/node/:id", h.UpdateNode)
		// видалити файл або папку
		storage.DELETE("/node/:id", h.DeleteNode)
		// загрузити фізичний файл у вже створену ноду
		storage.POST("/node/:id/upload", h.UploadFile)
		// скачати файл
		storage.GET("/node/:id/download", h.DownloadFile)		*/
	}
}
