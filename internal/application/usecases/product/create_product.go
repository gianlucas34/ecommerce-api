package usecases

import (
	ports "github.com/gianlucas34/ecommerce-api/internal/application/ports/repositories"
)

type CreateProduct struct {
	ProductRepository ports.ProductRepository
}

type CreateProductInput struct {
	Name  string
	Price float64
}

func NewCreateProduct(productRepository ports.ProductRepository) *CreateProduct {
	return &CreateProduct{
		ProductRepository: productRepository,
	}
}

func (c *CreateProduct) Execute(input CreateProductInput) error {
	return nil
}
