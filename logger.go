package logger

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var loggerType LogType

type LogType string
type LogLevel string

const (
	LogTypeConsole = "console"
	LogTypeJson    = "json"
)

const (
	LogLevelDisabled = "disabled"
	LogLevelDebug    = "debug"
	LogLevelInfo     = "info"
	LogLevelWarn     = "warn"
	LogLevelError    = "error"
)

type Logger struct {
	zerolog.Logger
}

func Init(logType LogType, logLevel LogLevel) {
	var level zerolog.Level
	switch logLevel {
	case LogLevelDisabled:
		level = zerolog.Disabled
	case LogLevelDebug:
		level = zerolog.DebugLevel
	case LogLevelWarn:
		level = zerolog.WarnLevel
	case LogLevelError:
		level = zerolog.ErrorLevel
	default:
		level = zerolog.InfoLevel
	}

	switch logType {
	case LogTypeJson:
		loggerType = logType
	default:
		loggerType = LogTypeConsole
	}

	zerolog.SetGlobalLevel(level)
}

func For(service string) Logger {
	var logger zerolog.Logger
	if loggerType == LogTypeJson {
		logger = log.With().Timestamp().Logger()
	} else {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	}

	return Logger{logger.With().Str("service", service).Logger()}
}

func (l Logger) Debug(ctx context.Context) *zerolog.Event {
	return l.Logger.Debug().Str(requestIDKey, getRequestIDFromContext(ctx))
}

func (l Logger) Info(ctx context.Context) *zerolog.Event {
	return l.Logger.Info().Str(requestIDKey, getRequestIDFromContext(ctx))
}

func (l Logger) Warn(ctx context.Context) *zerolog.Event {
	return l.Logger.Warn().Str(requestIDKey, getRequestIDFromContext(ctx))
}

func (l Logger) Error(ctx context.Context) *zerolog.Event {
	return l.Logger.Error().Str(requestIDKey, getRequestIDFromContext(ctx))
}
