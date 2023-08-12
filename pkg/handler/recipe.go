package handler

import (
	"wtfood/pkg/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckRecipe(c *gin.Context) {
	dishId := c.Param("dishId")
	ingredientId := c.Param("ingredientId")

	if dishId == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": model.ErrMissingDishId.Error()})
		return
	}

	if ingredientId == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": model.ErrMissingIngredientId.Error()})
		return
	}

	dish, ingredient, exist, err := h.Service.DishHasIngredient(dishId, ingredientId)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"dish": dish, "ingredient": ingredient, "exist": exist})
}
