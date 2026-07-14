package web

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func (h *Handler) Login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}
