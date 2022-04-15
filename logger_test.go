package log

import (
	"context"
	"testing"
)

func TestLoggerConsole(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, requestIDKey, "12345")

	Init(TypeConsole, LevelDebug)

	logger := For("test1")

	WithContext(ctx, logger).Debug().Msg("test debug")
	WithContext(ctx, logger).Info().Msg("test info")
	WithContext(ctx, logger).Warn().Msg("test warn")
	WithContext(ctx, logger).Error().Msg("test error")
}

func TestLoggerJSON(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, requestIDKey, "54321")

	Init(TypeJson, LevelInfo)

	logger := For("test2")

	WithContext(ctx, logger).Debug().Msg("test debug")
	WithContext(ctx, logger).Info().Msg("test info")
	WithContext(ctx, logger).Warn().Msg("test warn")
	WithContext(ctx, logger).Error().Msg("test error")
}
