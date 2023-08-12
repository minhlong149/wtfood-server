package middleware

import (
	"github.com/gin-gonic/gin"

	"wtfood/pkg/util"
)

func ErrorHandler(logger util.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			switch err.Err {

			default:
				logger.Error(err)
				c.JSON(500, gin.H{"error": "Something went wrong"})
				return
			}
		}
	}
}
