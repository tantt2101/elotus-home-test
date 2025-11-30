package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"github.com/joho/godotenv"
	"elotus-home-test/internal/api/router"
	"elotus-home-test/internal/config"
)

func main() {
	godotenv.Load()
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	mysqlDB, err := database.ConnectMySQL()
	if err != nil {
		fmt.Println("db fail:", err)
		return
	}

	router := router.NewRouter(mysqlDB)

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	fmt.Println("run with port:", port)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
