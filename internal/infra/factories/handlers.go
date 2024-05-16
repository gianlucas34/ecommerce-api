package factories

import (
	producthdls "github.com/gianlucas34/ecommerce-api/internal/infra/api/handlers/product"
	userhdls "github.com/gianlucas34/ecommerce-api/internal/infra/api/handlers/user"
)

type HandlersFactory struct {
	CreateUserHandler    *userhdls.CreateUserHandler
	CreateProductHandler *producthdls.CreateProductHandler
}

func NewHandlersFactory(usecases *UsecasesFactory) *HandlersFactory {
	return &HandlersFactory{
		CreateUserHandler:    userhdls.NewCreateUserHandler(usecases.CreateUser),
		CreateProductHandler: producthdls.NewCreateProductHandler(usecases.CreateProduct),
	}
}
