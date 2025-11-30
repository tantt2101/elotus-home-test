package services

import (
	"database/sql"
	"errors"
	"time"
	"elotus-home-test/internal/structs"
	"elotus-home-test/internal/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	DB *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{DB: db}
}

func (s *AuthService) Login(req structs.LoginRequest) (*structs.LoginResponse, error) {
	var id int64
	var hashedPassword string

	err := s.DB.QueryRow(
		"SELECT id, password FROM users WHERE username = ?",
		req.Username,
	).Scan(&id, &hashedPassword)

	if err == sql.ErrNoRows {
		return nil, errors.New("invalid username or password")
	}
	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)) != nil {
		return nil, errors.New("invalid username or password")
	}

	token, err := auth.GenerateToken(id, 60)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &structs.LoginResponse{
		AcceptToken:  token,
	}, nil
}

func (s *AuthService) RevokeToken(token string, expiresAt time.Time) error {
	_, err := s.DB.Exec(
		"INSERT INTO user_tokens (token, expires_at) VALUES (?, ?)",
		token,
		expiresAt,
	)
	return err
}

func (s *AuthService) CheckTokenRevoked(token string) (bool, error) {
	var exists bool
	err := s.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM user_tokens WHERE token = ?)",
		token,
	).Scan(&exists)

	return exists, err
}