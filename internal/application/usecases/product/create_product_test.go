package usecases

import (
	"testing"

	"github.com/gianlucas34/ecommerce-api/internal/errors"
	memoryrepos "github.com/gianlucas34/ecommerce-api/internal/infra/repositories/memory"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	t.Run("Should create product correctly", func(t *testing.T) {
		productRepository := memoryrepos.NewProductRepositoryMemory()
		createProduct := NewCreateProduct(productRepository)

		err := createProduct.Execute(CreateProductInput{
			Name:  "Produto",
			Price: 12.5,
		})

		assert.NoError(t, err)
		assert.Len(t, productRepository.Products, 1)
	})

	t.Run("Should return error if NewProduct() called with incorrect data", func(t *testing.T) {
		productRepository := memoryrepos.NewProductRepositoryMemory()
		createProduct := NewCreateProduct(productRepository)

		err := createProduct.Execute(CreateProductInput{
			Name:  "Produto",
			Price: 0,
		})

		assert.EqualError(t, err, errors.INVALID_PRODUCT_PRICE)
		assert.Len(t, productRepository.Products, 0)
	})

	t.Run("Should return error if ProductRepository.Create() fails", func(t *testing.T) {
		productRepository := memoryrepos.NewProductRepositoryMemory()
		createProduct := NewCreateProduct(productRepository)

		productRepository.FailCreate()

		err := createProduct.Execute(CreateProductInput{
			Name:  "Produto",
			Price: 12.5,
		})

		assert.EqualError(t, err, errors.CREATE_PRODUCT)
		assert.Len(t, productRepository.Products, 0)
	})
}
