package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func SampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("start sample middleware")
		c.Next()
		log.Println("end sample middleware")
	}
}
