package factories

import (
	"github.com/gianlucas34/ecommerce-api/internal/application/usecases/product"
	"github.com/gianlucas34/ecommerce-api/internal/application/usecases/user"
)

type UsecasesFactory struct {
	CreateUser    *user.CreateUserUsecase
	CreateProduct *product.CreateProductUsecase
}

func NewUsecasesFactory(repositories *RepositoriesFactory, adapters *AdaptersFactory) *UsecasesFactory {
	return &UsecasesFactory{
		CreateUser:    user.NewCreateUserUsecase(repositories.UserRepository, adapters.PasswordHashingAdapter),
		CreateProduct: product.NewCreateProductUsecase(repositories.ProductRepository),
	}
}
