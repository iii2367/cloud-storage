package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetTree(c *gin.Context) {

	userID := c.MustGet("userID").(uint)
	treeID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid UUID",
		})
		return
	}

	tree, err := h.service.GetTree(
		c.Request.Context(),
		treeID,
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
