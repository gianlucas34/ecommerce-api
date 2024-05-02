package mocks

import "github.com/gianlucas34/ecommerce-api/internal/errors"

type PasswordHashingMock struct {
	generateFailed bool
	compareFailed  bool
}

func NewPasswordHashingMock() *PasswordHashingMock {
	return &PasswordHashingMock{}
}

func (m *PasswordHashingMock) Generate(plaintext string) (string, error) {
	if m.generateFailed {
		return "", errors.NewInternalServerError(errors.GENERATE_HASH)
	}

	return "$2a$12$HR3X8o92WgUHOwnLeiNHaucRz/iEpKQ9X2rwQa/AZXrCscmTTYq6a", nil
}

func (m *PasswordHashingMock) Compare(plaintext string, digest string) (bool, error) {
	if m.compareFailed {
		return false, errors.NewInternalServerError(errors.COMPARE_PASSWORD)
	}

	return true, nil
}

func (m *PasswordHashingMock) FailGenerate() *PasswordHashingMock {
	m.generateFailed = true

	return m
}

func (m *PasswordHashingMock) FailCompare() *PasswordHashingMock {
	m.compareFailed = true

	return m
}
