package repositories

import entities "github.com/gianlucas34/ecommerce-api/internal/domain"

type ProductRepository interface {
	Create(product *entities.Product) error
	FindAll() ([]*entities.Product, error)
	Find(id string) (*entities.Product, error)
}
