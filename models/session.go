package models

import "database/sql"

type Session struct {
	ID     int
	UserID int
	// Token is only set when create a new session.
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	// 1. Create the session token
	// TODO: Implement SessionService.Create
	return nil, nil
}

func (ss *SessionService) User(token string) (*User, error) {

	return nil, nil
}
