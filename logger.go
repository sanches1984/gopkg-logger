package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	Verbose bool
	Jsonlog bool
	App     = "app"
)

type Level string

const (
	LevelDebug Level = "D"
	LevelInfo Level = "I"
	LevelWarning Level = "W"
	LevelError Level = "E"
	LevelFatal Level = "F"
)

var levelNameMap = map[Level]string {
	LevelDebug: "debug",
	LevelInfo: "info",
	LevelWarning: "warning",
	LevelError: "error",
	LevelFatal: "fatal",
}

type GetRequestIdFromContextGetterFn func(ctx context.Context) string

var GetRequestIdFromContextFn *GetRequestIdFromContextGetterFn

type jsonLog struct {
	Time     int64  `json:"time"`
	Request  string `json:"request"`
	Severity string `json:"severity"`
	Message  string `json:"message"`
}

func requestId(id interface{}) string {
	switch v := id.(type) {
	case string:
		return v
	case context.Context:
		if GetRequestIdFromContextFn != nil {
			return (*GetRequestIdFromContextFn)(v)
		}
	}
	return "<none>"
}

func jsonPrintf(severity string, ctx interface{}, format string, v ...interface{}) {
	bytes, err := json.Marshal(&jsonLog{time.Now().Unix(), requestId(ctx), severity, fmt.Sprintf(format, v...)})

	if err != nil {
		log.Printf("Logger failed to print message in JSON format")
		log.Printf("severity = %v", severity)
		log.Printf(format, v...)
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(bytes))
}

func Log(ctx interface{}, level Level, format string, v ...interface{}) {
	if Verbose != true && level == LevelDebug {
		return
	}
	if Jsonlog {
		jsonPrintf(levelNameMap[level], ctx, format, v...)
	} else {
		log.Printf("[%s] [%s] %s", level, requestId(ctx), fmt.Sprintf(format, v...))
	}
	if level == LevelFatal {
		os.Exit(1)
	}
}

func Debug(ctx interface{}, format string, v ...interface{}) {
	Log(ctx, LevelDebug, format, v...)
}

func Info(ctx interface{}, format string, v ...interface{}) {
	Log(ctx, LevelInfo, format, v...)
}

func Warning(ctx interface{}, format string, v ...interface{}) {
	Log(ctx, LevelWarning, format, v...)
}

func Error(ctx interface{}, format string, v ...interface{}) {
	Log(ctx, LevelError, format, v...)
}

func Fatal(ctx interface{}, format string, v ...interface{}) {
	Log(ctx, LevelFatal, format, v...)
}