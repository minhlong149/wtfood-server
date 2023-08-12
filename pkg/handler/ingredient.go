package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllIngredient(c *gin.Context) {
	ingredients, err := h.Service.GetAllIngredients()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"ingredients": ingredients})
}
