package product

import (
	"github.com/gianlucas34/ecommerce-api/internal/application/ports"
	"github.com/gianlucas34/ecommerce-api/internal/domain"
)

type CreateProductUsecase struct {
	ProductRepository ports.ProductRepository
}

type CreateProductUsecaseInput struct {
	Name  string
	Price float64
}

func NewCreateProductUsecase(productRepository ports.ProductRepository) *CreateProductUsecase {
	return &CreateProductUsecase{
		ProductRepository: productRepository,
	}
}

func (uc *CreateProductUsecase) Execute(input CreateProductUsecaseInput) error {
	product, err := domain.NewProduct(input.Name, input.Price)

	if err != nil {
		return err
	}

	err = uc.ProductRepository.Create(product)

	if err != nil {
		return err
	}

	return nil
}
