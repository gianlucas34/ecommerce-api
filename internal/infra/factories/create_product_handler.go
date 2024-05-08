package factories

import (
	usecases "github.com/gianlucas34/ecommerce-api/internal/application/usecases/product"
	handlers "github.com/gianlucas34/ecommerce-api/internal/infra/api/handlers/product"
	memoryrepos "github.com/gianlucas34/ecommerce-api/internal/infra/repositories/memory"
)

func CreateProductHandlerFactory() *handlers.CreateProductHandler {
	productRepository := memoryrepos.NewProductRepositoryMemory()
	createProductUsecase := usecases.NewCreateProduct(productRepository)
	createProductHandler := handlers.NewCreateProductHandler(createProductUsecase)

	return createProductHandler
}
