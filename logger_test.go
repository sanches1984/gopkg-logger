package log

import (
	"context"
	"github.com/go-chi/chi/middleware"
	"testing"
)

func TestLoggerConsole(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, middleware.RequestIDKey, "test")

	Init(TypeConsole, LevelDebug)

	logger := For("test1")

	WithContext(ctx, logger).Debug().Msg("test debug")
	WithContext(ctx, logger).Info().Msg("test info")
	WithContext(ctx, logger).Warn().Msg("test warn")
	WithContext(ctx, logger).Error().Msg("test error")
}

func TestLoggerJSON(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, middleware.RequestIDKey, "test")

	Init(TypeJson, LevelInfo)

	logger := For("test2")

	WithContext(ctx, logger).Debug().Msg("test debug")
	WithContext(ctx, logger).Info().Msg("test info")
	WithContext(ctx, logger).Warn().Msg("test warn")
	WithContext(ctx, logger).Error().Msg("test error")
}
