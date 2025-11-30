package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var Ctx = context.Background()

func ConnectMySQL() (*sql.DB, error) {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	db := os.Getenv("MYSQL_DATABASE")
	port := os.Getenv("MYSQL_PORT")

	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "3306"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		user, pass, host, port, db,
	)

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	log.Println("MySQL connected successfully!")

	return sqlDB, nil
}
