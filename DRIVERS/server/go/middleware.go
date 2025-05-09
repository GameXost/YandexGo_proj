// for http
package server

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

func JWTHTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte("your_jwt_secret"), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), "driver_id", claims["sub"])
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
}
