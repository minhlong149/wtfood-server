package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllDishes(c *gin.Context) {
	// TODO: Add pagination and search
	dishes, err := h.Service.GetAllDishes()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{
		"dishes": dishes,
		"total":  len(dishes),
	})
}

func (h *Handler) GetRandomDish(c *gin.Context) {
	dish, err := h.Service.GetRandomDish()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{
		"dish":        dish,
		"ingredients": len(dish.Ingredients),
	})
}

func (h *Handler) GetDishById(c *gin.Context) {
	dish, err := h.Service.GetDishById(c.Param("dishId"))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{
		"dish":        dish,
		"ingredients": len(dish.Ingredients),
	})
}

func (h *Handler) GetIngredientsByDishId(c *gin.Context) {
	dish, err := h.Service.GetDishById(c.Param("dishId"))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{
		"dish":        dish,
		"ingredients": dish.Ingredients,
	})
}
