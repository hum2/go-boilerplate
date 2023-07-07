package logger

import (
	"github.com/hum2/backend/internal/shared/logger"
	"github.com/rs/zerolog"
	"os"
)

type handler struct {
	logger zerolog.Logger
}

func New() logger.Logger {
	return &handler{
		logger: zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}
}

func (h *handler) Debug(msg string, fields ...map[string]interface{}) {
	h.logger.Debug().Fields(fields).Msg(msg)
}

func (h *handler) Info(msg string, fields ...map[string]interface{}) {
	h.logger.Info().Fields(fields).Msg(msg)
}

func (h *handler) Warn(msg string, fields ...map[string]interface{}) {
	h.logger.Warn().Fields(fields).Msg(msg)
}

func (h *handler) Error(msg string, fields ...map[string]interface{}) {
	h.logger.Error().Fields(fields).Msg(msg)
}

func (h *handler) Fatal(msg string, fields ...map[string]interface{}) {
	h.logger.Fatal().Fields(fields).Msg(msg)
}
