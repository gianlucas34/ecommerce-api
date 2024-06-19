package domain

import (
	"net/mail"
	"time"

	"github.com/gianlucas34/ecommerce-api/internal/errors"
	"github.com/google/uuid"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}

func NewUser(name string, email string, password string) (*User, error) {
	user := &User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
	err := user.Validate()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.NewValidationError(errors.EMPTY_USERNAME)
	}

	if u.Email == "" {
		return errors.NewValidationError(errors.EMPTY_USER_EMAIL)
	}

	_, err := mail.ParseAddress(u.Email)

	if err != nil {
		return errors.NewValidationError(errors.INVALID_USER_EMAIL)
	}

	if u.Password == "" {
		return errors.NewValidationError(errors.EMPTY_USER_PASSWORD)
	}

	return nil
}
