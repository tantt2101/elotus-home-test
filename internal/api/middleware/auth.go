package middleware

import (
	"context"
	"database/sql"
	"net/http"
	"strings"
	"fmt"
	"elotus-home-test/internal/auth"
	"elotus-home-test/internal/api/utils"
	"elotus-home-test/internal/services"
)

func AuthMiddleware(db *sql.DB) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            
            tokenHeader := r.Header.Get("Authorization")
            if !strings.HasPrefix(tokenHeader, "Bearer ") {
                utils.Error(w, "Missing or invalid token", http.StatusUnauthorized)
                return
            }

            tokenString := strings.TrimPrefix(tokenHeader, "Bearer ")

            claims, err := auth.ParseToken(tokenString)
            if err != nil {
                utils.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }

            authService := services.NewAuthService(db)
            revoked, err := authService.CheckTokenRevoked(tokenString)
            if err != nil {
                utils.Error(w, "Server error", http.StatusInternalServerError)
                return
            }

            if revoked {
                utils.Error(w, "Token revoked", http.StatusUnauthorized)
                return
            }
            ctx := context.WithValue(r.Context(), "user_id", fmt.Sprintf("%d", claims.UserID))

            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}