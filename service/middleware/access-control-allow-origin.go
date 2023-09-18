package middleware

import (
	"github.com/gin-gonic/gin"
)

func AccessPass() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
