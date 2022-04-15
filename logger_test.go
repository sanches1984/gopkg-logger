package logger

import (
	"context"
	"testing"
)

func TestLoggerConsole(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, requestIDKey, "12345")

	Init(LogTypeConsole, LogLevelDebug)

	logger := For("test1")

	logger.Debug(ctx).Msg("test debug")
	logger.Info(ctx).Msg("test info")
	logger.Warn(ctx).Msg("test warn")
	logger.Error(ctx).Msg("test error")
}

func TestLoggerJSON(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, requestIDKey, "54321")

	Init(LogTypeJson, LogLevelInfo)

	logger := For("test2")

	logger.Debug(ctx).Msg("test debug")
	logger.Info(ctx).Msg("test info")
	logger.Warn(ctx).Msg("test warn")
	logger.Error(ctx).Msg("test error")
}
