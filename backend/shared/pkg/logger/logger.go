// shared/pkg/logger/logger.go
package logger

import (
	"log/slog"
	"os"
)

type LogLevel string

const (
	LevelDebug LogLevel = "debug"
	LevelInfo  LogLevel = "info"
	LevelWarn  LogLevel = "warn"
	LevelError LogLevel = "error"
)

// Config конфигурация логгера
type Config struct {
	Level      LogLevel
	JSONFormat bool // true = JSON, false = Text
	AddSource  bool // Добавлять ли информацию о файле и строке
}

// New создает новый logger
func New(cfg Config) *slog.Logger {
	var level slog.Level

	switch cfg.Level {
	case LevelDebug:
		level = slog.LevelDebug
	case LevelInfo:
		level = slog.LevelInfo
	case LevelWarn:
		level = slog.LevelWarn
	case LevelError:
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: cfg.AddSource,
	}

	var handler slog.Handler
	if cfg.JSONFormat {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}

// NewDefault создает logger с дефолтными настройками для разработки
func NewDefault() *slog.Logger {
	return New(Config{
		Level:      LevelDebug,
		JSONFormat: false,
		AddSource:  true,
	})
}

// NewProduction создает logger для production
func NewProduction() *slog.Logger {
	return New(Config{
		Level:      LevelInfo,
		JSONFormat: true,
		AddSource:  false,
	})
}

// WithService добавляет service name к логгеру
func WithService(logger *slog.Logger, serviceName string) *slog.Logger {
	return logger.With(slog.String("service", serviceName))
}
