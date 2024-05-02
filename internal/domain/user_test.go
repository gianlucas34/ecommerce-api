package entities

import (
	"testing"
	"time"

	"github.com/gianlucas34/ecommerce-api/internal/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	t.Run("Should create user entity correctly", func(t *testing.T) {
		user, _ := NewUser("Usuário", "user@gmail.com", "senha")
		_, err := uuid.Parse(user.ID)

		if err != nil {
			t.Errorf(err.Error())
		}

		assert.Equal(t, "Usuário", user.Name)
		assert.Equal(t, "user@gmail.com", user.Email)
		assert.Equal(t, "senha", user.Password)

		_, err = time.Parse(time.RFC3339, user.CreatedAt)

		if err != nil {
			t.Errorf("A data gerada para o campo CreatedAt é inválida, forneça uma data do tipo RFC3339!")
		}

		_, err = time.Parse(time.RFC3339, user.UpdatedAt)

		if err != nil {
			t.Errorf("A data gerada para o campo UpdatedAt é inválida, forneça uma data do tipo RFC3339!")
		}
	})

	t.Run("Should return EMPTY_USERNAME exception if username is empty", func(t *testing.T) {
		_, err := NewUser("", "user@gmail.com", "senha")

		assert.EqualError(t, err, errors.EMPTY_USERNAME)
	})

	t.Run("Should return EMPTY_USER_EMAIL exception if email is empty", func(t *testing.T) {
		_, err := NewUser("Usuário", "", "senha")

		assert.EqualError(t, err, errors.EMPTY_USER_EMAIL)
	})

	t.Run("Should return INVALID_USER_EMAIL exception if email is invalid", func(t *testing.T) {
		_, err := NewUser("Usuário", "email", "senha")

		assert.EqualError(t, err, errors.INVALID_USER_EMAIL)
	})

	t.Run("Should return EMPTY_USER_PASSWORD exception if password is empty", func(t *testing.T) {
		_, err := NewUser("Usuário", "user@gmail.com", "")

		assert.EqualError(t, err, errors.EMPTY_USER_PASSWORD)
	})
}
