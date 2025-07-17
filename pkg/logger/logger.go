package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// New creates a new logger instance
func New() zerolog.Logger {
	// Set up pretty console logging for development
	if os.Getenv("GIN_MODE") != "release" {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	}

	// Set global log level
	level := zerolog.InfoLevel
	if os.Getenv("LOG_LEVEL") != "" {
		switch os.Getenv("LOG_LEVEL") {
		case "debug":
			level = zerolog.DebugLevel
		case "info":
			level = zerolog.InfoLevel
		case "warn":
			level = zerolog.WarnLevel
		case "error":
			level = zerolog.ErrorLevel
		case "fatal":
			level = zerolog.FatalLevel
		case "panic":
			level = zerolog.PanicLevel
		}
	}

	zerolog.SetGlobalLevel(level)

	// Add timestamp to all logs
	zerolog.TimeFieldFormat = time.RFC3339

	return log.Logger
}

// Logger returns the global logger instance
func Logger() zerolog.Logger {
	return log.Logger
} 