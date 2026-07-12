package http

import (
	"net/http"
	"cloud-storage/internal/auth"	
	"cloud-storage/internal/auth/dto"	
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service	*auth.Service
}

func NewHandler(service *auth.Service) *Handler {
	return &Handler {
		service: service,	
	}
}

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

func (h *Handler) Login(c *gin.Context) {
	var req dto.LoginRequest
	err := c.ShouldBindJSON(&req);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loginResponse, tokens, err := h.service.Login(
		c.Request.Context(),
		req.Email,
		req.Password,
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
    	Name:     "refresh_token",
	    Value:    tokens.RefreshToken,
    	Path:     "/",
	    HttpOnly: true,
    	Secure:   false, // true when https
	    SameSite: http.SameSiteLaxMode,
    	MaxAge:   60 * 60 * 24 * 30,
	})	
	c.JSON(http.StatusOK, loginResponse)
}
