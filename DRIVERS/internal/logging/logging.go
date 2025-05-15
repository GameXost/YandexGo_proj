package server

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

func LoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		resp, err := handler(ctx, req)
		log.Printf("Method: %s, Duration: %s, Error: %v", info.FullMethod, time.Since(start), err)
		return resp, err
	}
}

func LoggingHTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("HTTP %s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
