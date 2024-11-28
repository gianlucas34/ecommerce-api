package ports

import "github.com/gianlucas34/ecommerce-api/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindAll() ([]*domain.User, error)
	Find(id string) (*domain.User, error)
	FindByEmail(email string) (bool, error)
}
