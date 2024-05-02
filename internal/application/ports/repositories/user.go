package repositories

import entities "github.com/gianlucas34/ecommerce-api/internal/domain"

type UserRepository interface {
	Create(user *entities.User) error
	FindAll() ([]*entities.User, error)
	Find(id string) (*entities.User, error)
	FindByEmail(email string) (bool, error)
}
