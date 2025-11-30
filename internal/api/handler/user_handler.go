package handler

import (
	"encoding/json"
	"net/http"

	"elotus-home-test/internal/api/structs"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func RegisterUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req structs.RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			api.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if err := validate.Struct(req); err != nil {
			api.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userService := services.NewUserService(db)

		createdUser, err := userService.Register(req)
		if err != nil {
			api.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		api.Created(w, "Created successfully", createdUser)
	}
}
