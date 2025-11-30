package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"elotus-home-test/internal/auth"
	"elotus-home-test/internal/api/utils"
	"elotus-home-test/internal/services"
	"elotus-home-test/internal/structs"
)

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			utils.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var req structs.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if err := utils.Validate.Struct(req); err != nil {
			utils.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		authService := services.NewAuthService(db)
		resp, err := authService.Login(req)
		if err != nil {
			utils.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		utils.Success(w, "Login success", resp)
	}
}

func Logout(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			utils.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		tokenHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(tokenHeader, "Bearer ") {
			utils.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(tokenHeader, "Bearer ")

		claims, err := auth.ParseToken(tokenStr)
		if err != nil {
			utils.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		expiresAt := claims.ExpiresAt.Time
		if expiresAt.Before(time.Now()) {
			utils.Error(w, "Token already expired", http.StatusBadRequest)
			return
		}

		authService := services.NewAuthService(db)
		if err := authService.RevokeToken(tokenStr, expiresAt); err != nil {
			utils.Error(w, "Failed to revoke token", http.StatusInternalServerError)
			return
		}

		utils.Success(w, "Logout successfully", nil)
	}
}
