package factories

import (
	productucs "github.com/gianlucas34/ecommerce-api/internal/application/usecases/product"
	userucs "github.com/gianlucas34/ecommerce-api/internal/application/usecases/user"
)

type UsecasesFactory struct {
	CreateUser    *userucs.CreateUser
	CreateProduct *productucs.CreateProduct
}

func NewUsecasesFactory(repositories *RepositoriesFactory, adapters *AdaptersFactory) *UsecasesFactory {
	return &UsecasesFactory{
		CreateUser:    userucs.NewCreateUser(repositories.UserRepository, adapters.PasswordHashingAdapter),
		CreateProduct: productucs.NewCreateProduct(repositories.ProductRepository),
	}
}
