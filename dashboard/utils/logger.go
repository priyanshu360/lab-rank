package utils

import (
	"log/slog"
	"os"

	"github.com/priyanshu360/lab-rank/dashboard/config"
)

// NewLogger creates a new logger with the specified configuration.
func NewLogger(config config.LoggerConfig) *slog.Logger {
	logLevel := slog.LevelInfo // Default log level
	if config.LogLevel == "ERROR" {
		logLevel = slog.LevelError
	} else if config.LogLevel == "WARNING" {
		logLevel = slog.LevelWarn
	} else if config.LogLevel == "DEBUG" {
		logLevel = slog.LevelDebug
	}

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     logLevel,
		AddSource: true,
	})
	logger := slog.New(handler)
	return logger
}
