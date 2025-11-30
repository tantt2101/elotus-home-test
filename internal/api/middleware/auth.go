package middleware

import (
	"context"
	"net/http"
	"strings"

	"elotus-home-test/internal/api/utils"
	"elotus-home-test/internal/api/services"
	"elotus-home-test/internal/api/utils"
)

func AuthMiddleware(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenHeader := r.Header.Get("Authorization")
			if !strings.HasPrefix(tokenHeader, "Bearer ") {
				utils.Error(w, "Missing or invalid token", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(auth, "Bearer ")

			claims, err := auth.ParseToken(tokenString)
			if err != nil {
				utils.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			tokenService := services.NewTokenService(db)
			revoked, err := tokenService.IsRevoked(tokenString)
			if err != nil {
				utils.Error(w, "Server error", http.StatusInternalServerError)
				return
			}

			if revoked {
				utils.Error(w, "Token revoked", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
