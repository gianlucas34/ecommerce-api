package domain

import (
	"time"

	"github.com/gianlucas34/ecommerce-api/internal/errors"
	"github.com/google/uuid"
)

type Product struct {
	ID        string
	Name      string
	Price     float64
	CreatedAt string
	UpdatedAt string
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        uuid.New().String(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
	err := product.Validate()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.NewValidationError(errors.EMPTY_PRODUCT_NAME)
	}

	if p.Price <= 0 {
		return errors.NewValidationError(errors.INVALID_PRODUCT_PRICE)
	}

	return nil
}
