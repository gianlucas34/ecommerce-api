package userucs

import (
	"github.com/gianlucas34/ecommerce-api/internal/application/ports/repositories"
	"github.com/gianlucas34/ecommerce-api/internal/application/ports/security"
	"github.com/gianlucas34/ecommerce-api/internal/domain"
	"github.com/gianlucas34/ecommerce-api/internal/errors"
)

type CreateUser struct {
	UserRepository  repositories.UserRepository
	PasswordHashing security.PasswordHashing
}

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
}

func NewCreateUser(userRepository repositories.UserRepository, passwordHashing security.PasswordHashing) *CreateUser {
	return &CreateUser{
		UserRepository:  userRepository,
		PasswordHashing: passwordHashing,
	}
}

func (c *CreateUser) Execute(input CreateUserInput) error {
	passwordHash, err := c.PasswordHashing.Generate(input.Password)

	if err != nil {
		return err
	}

	user, err := domain.NewUser(input.Name, input.Email, passwordHash)

	if err != nil {
		return err
	}

	alreadyInUse, err := c.UserRepository.FindByEmail(user.Email)

	if err != nil {
		return err
	}

	if alreadyInUse {
		return errors.NewFieldInUseError(errors.ALREADY_IN_USE_EMAIL)
	}

	err = c.UserRepository.Create(user)

	if err != nil {
		return err
	}

	return nil
}
