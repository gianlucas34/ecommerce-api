package factories

import (
	"github.com/gianlucas34/ecommerce-api/internal/application/ports"
	"github.com/gianlucas34/ecommerce-api/internal/mocks"
)

type AdaptersFactory struct {
	PasswordHashingAdapter ports.PasswordHashing
}

func NewAdaptersFactory() *AdaptersFactory {
	return &AdaptersFactory{
		PasswordHashingAdapter: mocks.NewPasswordHashingMock(),
	}
}
