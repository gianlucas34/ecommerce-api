package factories

import (
	"github.com/gianlucas34/ecommerce-api/internal/application/ports/security"
	"github.com/gianlucas34/ecommerce-api/internal/mocks"
)

type AdaptersFactory struct {
	PasswordHashingAdapter security.PasswordHashing
}

func NewAdaptersFactory() *AdaptersFactory {
	return &AdaptersFactory{
		PasswordHashingAdapter: mocks.NewPasswordHashingMock(),
	}
}
