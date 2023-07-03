package models

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	DefaultResetDuration = 1 * time.Hour
)

type PasswordReset struct {
	ID     int
	UserID int
	// Token is only set when a Password reset
	Token     string
	TokenHash string
	ExpiresAt time.Time
}

type PasswordResetService struct {
	DB *sql.DB
	// BytesPerToken is used to determine how many bytes are in a token.

	BytesPerToken int
	// Duration is the amout of time that a PasswordReset is valid for.
	// Default is 1 hour.
	Duration time.Duration
}

func (service *PasswordResetService) Create(email string) (*PasswordReset, error) {
	return nil, fmt.Errorf("TODO: Implement PasswordResetService.Create")
}

func (sercie *PasswordResetService) Consume(token string) (*User, error) {
	return nil, fmt.Errorf("TODO: Implement PasswordResetService.Consume")
}
