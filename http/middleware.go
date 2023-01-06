package http

import (
	"mck-p/modi/log"

	"github.com/gin-gonic/gin"
)

func SetRequestVariables(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", logger)

		c.Next()
	}
}

func SetGlobalHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Server", "Modi")

		c.Next()
	}
}
