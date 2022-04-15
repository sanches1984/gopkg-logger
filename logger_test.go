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

	logger.Debug(ctx).Msg("test debug")
	logger.Info(ctx).Msg("test info")
	logger.Warn(ctx).Msg("test warn")
	logger.Error(ctx).Msg("test error")
}

func TestLoggerJSON(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, requestIDKey, "54321")

	Init(TypeJson, LevelInfo)

	logger := For("test2")

	logger.Debug(ctx).Msg("test debug")
	logger.Info(ctx).Msg("test info")
	logger.Warn(ctx).Msg("test warn")
	logger.Error(ctx).Msg("test error")
}
