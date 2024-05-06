package usecases

import (
	ports "github.com/gianlucas34/ecommerce-api/internal/application/ports/repositories"
	entities "github.com/gianlucas34/ecommerce-api/internal/domain"
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
	product, err := entities.NewProduct(input.Name, input.Price)

	if err != nil {
		return err
	}

	err = c.ProductRepository.Create(product)

	if err != nil {
		return err
	}

	return nil
}
