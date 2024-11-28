package user

import (
	"testing"

	"github.com/gianlucas34/ecommerce-api/internal/errors"
	"github.com/gianlucas34/ecommerce-api/internal/infra/repositories/memory"
	"github.com/gianlucas34/ecommerce-api/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserUsecase(t *testing.T) {
	t.Run("Should create user correctly", func(t *testing.T) {
		userRepository := memory.NewUserRepository()
		passwordHashing := mocks.NewPasswordHashingMock()
		createUserUsecase := NewCreateUserUsecase(userRepository, passwordHashing)

		err := createUserUsecase.Execute(CreateUserUsecaseInput{
			Name:     "Usuário",
			Email:    "user@gmail.com",
			Password: "senha",
		})

		assert.NoError(t, err)
		assert.Len(t, userRepository.Users, 1)
	})

	t.Run("Should return error if PasswordHashing.Generate() fails", func(t *testing.T) {
		userRepository := memory.NewUserRepository()
		passwordHashing := mocks.NewPasswordHashingMock()
		createUserUsecase := NewCreateUserUsecase(userRepository, passwordHashing)

		passwordHashing.FailGenerate()

		err := createUserUsecase.Execute(CreateUserUsecaseInput{
			Name:     "Usuário",
			Email:    "user@gmail.com",
			Password: "senha",
		})

		assert.EqualError(t, err, errors.GENERATE_HASH)
		assert.Len(t, userRepository.Users, 0)
	})

	t.Run("Should return error if NewUser() called with incorrect data", func(t *testing.T) {
		userRepository := memory.NewUserRepository()
		passwordHashing := mocks.NewPasswordHashingMock()
		createUserUsecase := NewCreateUserUsecase(userRepository, passwordHashing)

		err := createUserUsecase.Execute(CreateUserUsecaseInput{
			Name:     "",
			Email:    "user@gmail.com",
			Password: "senha",
		})

		assert.EqualError(t, err, errors.EMPTY_USERNAME)
		assert.Len(t, userRepository.Users, 0)
	})

	t.Run("Should return error if UserRepository.FindByEmail() fails", func(t *testing.T) {
		userRepository := memory.NewUserRepository()
		passwordHashing := mocks.NewPasswordHashingMock()
		createUserUsecase := NewCreateUserUsecase(userRepository, passwordHashing)

		userRepository.FailFindByEmail()

		err := createUserUsecase.Execute(CreateUserUsecaseInput{
			Name:     "Usuário",
			Email:    "user@gmail.com",
			Password: "senha",
		})

		assert.EqualError(t, err, errors.FIND_USER_BY_EMAIL)
		assert.Len(t, userRepository.Users, 0)
	})

	t.Run("Should return ALREADY_IN_USE_EMAIL exception if email already in use", func(t *testing.T) {
		userRepository := memory.NewUserRepository()
		passwordHashing := mocks.NewPasswordHashingMock()
		createUserUsecase := NewCreateUserUsecase(userRepository, passwordHashing)

		userRepository.EmailAlreadyInUse()

		err := createUserUsecase.Execute(CreateUserUsecaseInput{
			Name:     "Usuário",
			Email:    "user@gmail.com",
			Password: "senha",
		})

		assert.EqualError(t, err, errors.ALREADY_IN_USE_EMAIL)
		assert.Len(t, userRepository.Users, 0)
	})

	t.Run("Should return error if UserRepository.Create() fails", func(t *testing.T) {
		userRepository := memory.NewUserRepository()
		passwordHashing := mocks.NewPasswordHashingMock()
		createUserUsecase := NewCreateUserUsecase(userRepository, passwordHashing)

		userRepository.FailCreate()

		err := createUserUsecase.Execute(CreateUserUsecaseInput{
			Name:     "Usuário",
			Email:    "user@gmail.com",
			Password: "senha",
		})

		assert.EqualError(t, err, errors.CREATE_USER)
		assert.Len(t, userRepository.Users, 0)
	})
}
