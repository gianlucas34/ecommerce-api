package factories

import (
	"github.com/gianlucas34/ecommerce-api/internal/application/ports"
	memoryrepos "github.com/gianlucas34/ecommerce-api/internal/infra/repositories/memory"
)

type RepositoriesFactory struct {
	UserRepository    ports.UserRepository
	ProductRepository ports.ProductRepository
}

func NewRepositoriesFactory() *RepositoriesFactory {
	return &RepositoriesFactory{
		UserRepository:    memoryrepos.NewUserRepository(),
		ProductRepository: memoryrepos.NewProductRepository(),
	}
}
