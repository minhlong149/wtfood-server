package handler

// import (
// 	"github.com/gin-gonic/gin"

// 	"wtfood/pkg/models"
// )

// type RecipeHandler struct {
// 	RecipeService models.RecipeService
// }

// func (h *RecipeHandler) CheckRecipe(c *gin.Context) {
// 	dish := c.Param("dish")
// 	ingredient := c.Param("ingredient")
// 	if dish == "" || ingredient == "" {
// 		c.JSON(400, gin.H{"error": "dish or ingredient is empty"})
// 		return
// 	}

// 	ingredients, err := h.RecipeService.CheckRecipe(dish, ingredient)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(200, ingredients)
// }
