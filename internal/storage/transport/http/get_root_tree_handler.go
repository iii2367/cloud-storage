package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRootTree(c *gin.Context) {

	userID := c.MustGet("userID").(uint)

	tree, err := h.service.GetRootTree(
		c.Request.Context(),
		userID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, tree)
}
