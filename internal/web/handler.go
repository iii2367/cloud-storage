package web

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func (h *Handler) Signup(c *gin.Context) {
	c.HTML(200, "signup.html", gin.H{})
}

func (h *Handler) Storage(c *gin.Context) {
	c.HTML(200, "storage.html", gin.H{})
}
