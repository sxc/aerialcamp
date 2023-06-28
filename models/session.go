package models

import (
	"database/sql"
	"fmt"

	"github.com/sxc/aerialcamp/rand"
)

const (
	MinBytesPerToken = 32
)

type Session struct {
	ID     int
	UserID int
	// Token is only set when create a new session.
	Token     string
	TokenHash string
}

type SessionService struct {
	DB            *sql.DB
	BytesPerToken int
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	// 1. Create the session token
	// TODO: Implement SessionService.Create
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	session := Session{
		UserID: userID,
		Token:  token,
	}
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {

	return nil, nil
}
