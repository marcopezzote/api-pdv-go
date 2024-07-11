package middleware

import (
	"api-pdv-golang/config"
	"api-pdv-golang/handlers"
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type contextKey string

const userContextKey contextKey = "user"

func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization header is required", http.StatusUnauthorized)
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        claims := &handlers.Claims{}

        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return config.JwtKey, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        ctx := context.WithValue(r.Context(), userContextKey, claims.Email)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
