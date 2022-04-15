package log

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var loggerType Type

type Type string
type Level string

const (
	TypeConsole = "console"
	TypeJson    = "json"
)

const (
	LevelDisabled = "disabled"
	LevelDebug    = "debug"
	LevelInfo     = "info"
	LevelWarn     = "warn"
	LevelError    = "error"
)

type Logger struct {
	zerolog.Logger
}

func Init(logType Type, logLevel Level) {
	var level zerolog.Level
	switch logLevel {
	case LevelDisabled:
		level = zerolog.Disabled
	case LevelDebug:
		level = zerolog.DebugLevel
	case LevelWarn:
		level = zerolog.WarnLevel
	case LevelError:
		level = zerolog.ErrorLevel
	default:
		level = zerolog.InfoLevel
	}

	switch logType {
	case TypeJson:
		loggerType = logType
	default:
		loggerType = TypeConsole
	}

	zerolog.SetGlobalLevel(level)
}

func For(service string) *zerolog.Logger {
	var logger zerolog.Logger
	if loggerType == TypeJson {
		logger = log.With().Timestamp().Logger()
	} else {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	}

	logger = logger.With().Str("service", service).Logger()
	return &logger
}

func WithContext(ctx context.Context, logger *zerolog.Logger) *zerolog.Logger {
	l := logger.With().Str(requestIDKey, getRequestIDFromContext(ctx)).Logger()
	return &l
}
