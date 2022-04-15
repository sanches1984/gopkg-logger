package log

import (
	"context"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"net/http"
)

func NewLogServerInterceptor(logger zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		WithContext(ctx, logger).Debug().Str("method", info.FullMethod).Msg("grpc unary call")
		resp, err := handler(ctx, req)
		return resp, err
	}
}

func NewLogServerMiddleware(logger zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			WithContext(r.Context(), logger).Debug().Str("method", r.Method).Str("uri", r.RequestURI).Msg("http handler call")
			next.ServeHTTP(w, r)
		})
	}
}
