package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *Handler, m *Middleware) {
	storage := r.Group("/api/storage")
	storage.Use(m.AuthMiddleware())
	{
		// дерево
		storage.GET("/tree", h.GetRootTree) // дає наслідників корня
		storage.GET("/tree/:id", h.GetTree) // дає наслідників відносно id папки

		// одна нода
		//		storage.GET("/nodes/:id", h.GetNode) // дає ноду за id

		// папки
		storage.POST("/folders", h.CreateFolder) // створює ноду папки
		storage.POST("/root", h.CreateRoot)      // створює початкову папку користувача

		// файли
		storage.POST("/files", h.UploadFile) // створює ноду файла та закріплює за нею файл

		// зміна
		//		storage.PATCH("/nodes/:id", h.UpdateNode) // оновлює дані доди які надасть користувач

		// видалення
		//		storage.DELETE("/nodes", h.DeleteNodes)		// видаля всі ноди користувача
		storage.DELETE("/nodes/:id", h.DeleteNode) // видаля конкретну ноду та її наслідників якщо файл

		// скачування
		storage.GET("/files/:id/download", h.DownloadFile) // повертає файл для завантаження
	}
}
