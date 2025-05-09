package server

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

func AuthInterceptor(publicKey []byte) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		// public method skip  jwt not required
		if info.FullMethod == "/drivers.Drivers/PublicMethod" {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
		}

		tokens := md.Get("authorization")
		if len(tokens) == 0 {
			return nil, status.Error(codes.Unauthenticated, "authorization token is not provided")
		}

		tokenString := strings.TrimPrefix(tokens[0], "Bearer ")
		driverID, err := validateToken(tokenString, publicKey)
		if err != nil {
			return nil, status.Error(codes.PermissionDenied, "permission denied due to invalid token")
		}
		ctx = context.WithValue(ctx, "driver_id", driverID)
		return handler(ctx, req)
	}
}

func validateToken(tokenString string, publicKey []byte) (string, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwt.ParseRSAPublicKeyFromPEM(publicKey)
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
