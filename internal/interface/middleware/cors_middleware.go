package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hum2/backend/internal/interface/config"
	"strings"
	"time"
)

// CorsMiddleware CORS middleware
func CorsMiddleware(config *config.Config) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: strings.Split(config.CORS.AllowOrigins, ","),
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		MaxAge: 24 * time.Hour,
	})
}
