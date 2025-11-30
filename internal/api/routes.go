package api

import (
	"database/sql"
	"net/http"
	"elotus-home-test/internal/api/handler"
)

func NewRouter(mysqlDB *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }
        handler.PingHandler(w, r)
    })

	return mux
}
