package factories

import (
	"github.com/gianlucas34/ecommerce-api/internal/application/ports/repositories"
	memoryrepos "github.com/gianlucas34/ecommerce-api/internal/infra/repositories/memory"
)

type RepositoriesFactory struct {
	UserRepository    repositories.UserRepository
	ProductRepository repositories.ProductRepository
}

func NewRepositoriesFactory() *RepositoriesFactory {
	return &RepositoriesFactory{
		UserRepository:    memoryrepos.NewUserRepositoryMemory(),
		ProductRepository: memoryrepos.NewProductRepositoryMemory(),
	}
}
