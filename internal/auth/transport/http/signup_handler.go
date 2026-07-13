package http

import (
	"cloud-storage/internal/auth/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Signup(c *gin.Context) {
	var req dto.SignupRequest
	err := c.ShouldBindJSON(&req);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	signupResponse, err := h.service.Signup(
		c.Request.Context(),
		req.Name,
		req.Email,
		req.Password,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, signupResponse)
}
