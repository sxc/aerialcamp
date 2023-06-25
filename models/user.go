package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Email        string
	Passwordhash string
}

type UserService struct {
	DB *sql.DB
}

// type NewUser struct {
// 	Email    string
// 	Password string
// }

func (us *UserService) Create(email, password string) (*User, error) {
	email = strings.ToLower(email)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("could not hash password: %v", err)
	}

	passwordHash := string(hashedBytes)

	user := User{
		Email:        email,
		Passwordhash: passwordHash,
	}

	row := us.DB.QueryRow(
		`INSERT INTO users (email, password_hash) 
		VALUES ($1, $2) RETURNING id`, email, passwordHash)
	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("could not create user: %v", err)
	}
	return &user, nil
}
