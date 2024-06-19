package productucs

import (
	"github.com/gianlucas34/ecommerce-api/internal/application/ports/repositories"
	"github.com/gianlucas34/ecommerce-api/internal/domain"
)

type CreateProduct struct {
	ProductRepository repositories.ProductRepository
}

type CreateProductInput struct {
	Name  string
	Price float64
}

func NewCreateProduct(productRepository repositories.ProductRepository) *CreateProduct {
	return &CreateProduct{
		ProductRepository: productRepository,
	}
}

func (c *CreateProduct) Execute(input CreateProductInput) error {
	product, err := domain.NewProduct(input.Name, input.Price)

	if err != nil {
		return err
	}

	err = c.ProductRepository.Create(product)

	if err != nil {
		return err
	}

	return nil
}
