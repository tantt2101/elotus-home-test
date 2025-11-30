package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"elotus-home-test/internal/api/utils"
	"elotus-home-test/internal/structs"
	"elotus-home-test/internal/services"
)

func RegisterUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req structs.RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if err := utils.Validate.Struct(req); err != nil {
			utils.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userService := services.NewUserService(db)

		createdUser, err := userService.RegisterUser(req)
		if err != nil {
			utils.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		utils.Created(w, "Created successfully", createdUser)
	}
}
