package factories

import "github.com/gianlucas34/ecommerce-api/internal/mocks"

type AdaptersFactory struct {
	PasswordHashingAdapter *mocks.PasswordHashingMock
}

func NewAdaptersFactory() *AdaptersFactory {
	return &AdaptersFactory{
		PasswordHashingAdapter: mocks.NewPasswordHashingMock(),
	}
}
