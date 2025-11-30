package router

import (
	"database/sql"
	"net/http"
	"elotus-home-test/internal/api/handler"
)

func NewRouter(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }
        handler.RegisterUser(db)(w, r)
    })
	mux.HandleFunc("/api/auth/login", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }
        handler.Login(db)(w, r)
    })
	mux.HandleFunc("/api/auth/logout", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }
        handler.Logout(db)(w, r)
    })

	return mux
}
