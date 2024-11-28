package user

import (
	"github.com/gianlucas34/ecommerce-api/internal/application/ports"
	"github.com/gianlucas34/ecommerce-api/internal/domain"
	"github.com/gianlucas34/ecommerce-api/internal/errors"
)

type CreateUserUsecase struct {
	UserRepository  ports.UserRepository
	PasswordHashing ports.PasswordHashing
}

type CreateUserUsecaseInput struct {
	Name     string
	Email    string
	Password string
}

func NewCreateUserUsecase(userRepository ports.UserRepository, passwordHashing ports.PasswordHashing) *CreateUserUsecase {
	return &CreateUserUsecase{
		UserRepository:  userRepository,
		PasswordHashing: passwordHashing,
	}
}

func (uc *CreateUserUsecase) Execute(input CreateUserUsecaseInput) error {
	passwordHash, err := uc.PasswordHashing.Generate(input.Password)

	if err != nil {
		return err
	}

	user, err := domain.NewUser(input.Name, input.Email, passwordHash)

	if err != nil {
		return err
	}

	alreadyInUse, err := uc.UserRepository.FindByEmail(user.Email)

	if err != nil {
		return err
	}

	if alreadyInUse {
		return errors.NewFieldInUseError(errors.ALREADY_IN_USE_EMAIL)
	}

	err = uc.UserRepository.Create(user)

	if err != nil {
		return err
	}

	return nil
}
