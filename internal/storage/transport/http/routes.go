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
		storage.POST("/nodes", h.CreateNode)
		// змінити назву/опис
		storage.PATCH("/node/:id", h.UpdateNode)
		// видалити файл або папку
		storage.DELETE("/node/:id", h.DeleteNode)
		// загрузити фізичний файл у вже створену ноду
		storage.POST("/node/:id/upload", h.UploadFile)
		// скачати файл
		storage.GET("/node/:id/download", h.DownloadFile)		*/

		// дерево
//		storage.GET("/tree", h.GetTree) 	// дає наслідників корня
//		storage.GET("/tree/:id", h.GetTree) // дає наслідників відносно id папки

		// одна нода
//		storage.GET("/nodes/:id", h.GetNode) // дає ноду за id

		// папки
		storage.POST("/folders", h.CreateFolder) // створює ноду папки

		// файли
//		storage.POST("/files", h.UploadFile) // створює ноду файла та закріплює за нею файл

		// зміна
//		storage.PATCH("/nodes/:id", h.UpdateNode) // оновлює дані доди які надасть користувач

		// видалення
//		storage.DELETE("/nodes/:id", h.DeleteNode) // видаля конкретну ноду та її наслідників якщо файл

		// скачування
//		storage.GET("/files/:id/download", h.DownloadFile) // повертає файл для завантаження
	}
}
