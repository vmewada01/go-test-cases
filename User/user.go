package user

import (
	"context"
	"errors"
	"strings"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

func AddUser(ctx context.Context, u *User) (*User, error) {
	// Basic validation of the user struct fields
	if strings.TrimSpace(u.Name) == "" {
		return nil, errors.New("name cannot be empty")
	}
	if strings.TrimSpace(u.Email) == "" || !strings.Contains(u.Email, "@") {
		return nil, errors.New("invalid email")
	}
	if u.Age < 0 {
		return nil, errors.New("age cannot be negative")
	}
	if len(u.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	// Return the user object without any database operation
	return u, nil

}
