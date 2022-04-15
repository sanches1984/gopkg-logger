package log

import (
	"context"
)

const requestIDKey = "request-id"

func getRequestIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if reqID, ok := ctx.Value(requestIDKey).(string); ok {
		return reqID
	}
	return ""
}
