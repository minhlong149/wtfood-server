package middleware

import (
	"github.com/gin-gonic/gin"

	"wtfood/pkg/model"
	"wtfood/pkg/util"
)

func ErrorHandler(logger util.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			switch err.Err {

			case model.ErrWrongCredentials:
				c.JSON(401, gin.H{"error": err.Error()})

			case model.ErrDishNotFound:
				c.JSON(404, gin.H{"error": err.Error()})

			case model.ErrIngredientNotFound:
				c.JSON(404, gin.H{"error": err.Error()})

			default:
				logger.Error(err)
				c.JSON(500, gin.H{"error": "Something went wrong"})
				return
			}
		}
	}
}
