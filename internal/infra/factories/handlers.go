package factories

import (
	"github.com/gianlucas34/ecommerce-api/internal/infra/api/handlers/product"
	"github.com/gianlucas34/ecommerce-api/internal/infra/api/handlers/user"
)

type HandlersFactory struct {
	CreateUserHandler    *user.CreateUserHandler
	CreateProductHandler *product.CreateProductHandler
}

func NewHandlersFactory(usecases *UsecasesFactory) *HandlersFactory {
	return &HandlersFactory{
		CreateUserHandler:    user.NewCreateUserHandler(usecases.CreateUser),
		CreateProductHandler: product.NewCreateProductHandler(usecases.CreateProduct),
	}
}
