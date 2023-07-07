package middleware

import (
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"golang.org/x/exp/slog"
	"os"
)

// LoggerMiddleware logging middleware
func LoggerMiddleware() gin.HandlerFunc {
	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}),
	)
	return sloggin.New(logger)
}
