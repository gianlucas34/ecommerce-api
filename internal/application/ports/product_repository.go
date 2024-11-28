package ports

import "github.com/gianlucas34/ecommerce-api/internal/domain"

type ProductRepository interface {
	Create(product *domain.Product) error
	FindAll() ([]*domain.Product, error)
	Find(id string) (*domain.Product, error)
}
