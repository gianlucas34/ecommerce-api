package factories

import memoryrepos "github.com/gianlucas34/ecommerce-api/internal/infra/repositories/memory"

type RepositoriesFactory struct {
	UserRepository    *memoryrepos.UserRepositoryMemory
	ProductRepository *memoryrepos.ProductRepositoryMemory
}

func NewRepositoriesFactory() *RepositoriesFactory {
	return &RepositoriesFactory{
		UserRepository:    memoryrepos.NewUserRepositoryMemory(),
		ProductRepository: memoryrepos.NewProductRepositoryMemory(),
	}
}
