package server

import (
	"context"
	"crypto/rsa"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

const DriverIDKey contextKey = "driver_id"

func AuthInterceptor(publicKey *rsa.PublicKey, disableAuth bool) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if disableAuth {
			return handler(ctx, req)
		}

		if isPublicMethod(info.FullMethod) {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		tokens := md.Get("authorization")
		if len(tokens) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing token")
		}

		driverID, err := validateToken(strings.TrimPrefix(tokens[0], "Bearer "), publicKey)
		if err != nil {
			return nil, status.Errorf(codes.PermissionDenied, "invalid token: %v", err)
		}

		return handler(context.WithValue(ctx, DriverIDKey, driverID), req)
	}
}

func JWTHTTPMiddleware(publicKey *rsa.PublicKey) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if isPublicRoute(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}

			tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
			driverID, err := validateToken(tokenString, publicKey)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), DriverIDKey, driverID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func isPublicMethod(method string) bool {
	publicMethods := map[string]bool{
		"none shit here": true,
	}
	return publicMethods[method]
}

func isPublicRoute(path string) bool {
	publicRoutes := map[string]bool{
		"shit": true,
	}
	return publicRoutes[path]
}

func validateToken(tokenString string, publicKey *rsa.PublicKey) (string, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return publicKey, nil
		},
		jwt.WithValidMethods([]string{"RS256"}),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return "", fmt.Errorf("token is not valid %w", err)
	}
	if !token.Valid {
		return "", fmt.Errorf("token is not valid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", fmt.Errorf("token is not valid format")
	}

	exp, err := claims.GetExpirationTime()
	if err != nil {
		return "", fmt.Errorf("invalid token expired: %w", err)
	}

	if exp.Time.Before(time.Now().UTC()) {
		return "", fmt.Errorf("token expired")
	}

	driverID, ok := claims["sub"].(string)
	if !ok || driverID == "" {
		return "", fmt.Errorf("invalid token or miss sub claim")
	}
	return driverID, nil

}
