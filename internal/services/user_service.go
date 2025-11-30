package services

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"

	"elotus-home-test/internal/api/structs"
)

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) RegisterUser(req structs.RegisterRequest) (*structs.UserResponse, error) {
	var exists bool
	err := s.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", req.Username).Scan(&exists)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	result, err := s.DB.Exec(
		"INSERT INTO users (username, password) VALUES (?, ?)",
		req.Username,
		string(hashedPassword),
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user := &structs.UserResponse{
		ID:       id,
		Username: req.Username,
	}

	return user, nil
}
