package memoryrepos

import (
	entities "github.com/gianlucas34/ecommerce-api/internal/domain"
	"github.com/gianlucas34/ecommerce-api/internal/errors"
)

type ProductRepositoryMemory struct {
	Products      []*entities.Product
	createFailed  bool
	findAllFailed bool
	findFailed    bool
}

func NewProductRepositoryMemory() *ProductRepositoryMemory {
	return &ProductRepositoryMemory{
		Products: make([]*entities.Product, 0),
	}
}

func (r *ProductRepositoryMemory) Create(product *entities.Product) error {
	if r.createFailed {
		return errors.NewInternalServerError(errors.CREATE_PRODUCT)
	}

	r.Products = append(r.Products, product)

	return nil
}

func (r *ProductRepositoryMemory) FindAll() ([]*entities.Product, error) {
	if r.findAllFailed {
		return nil, errors.NewInternalServerError(errors.FIND_ALL_PRODUCTS)
	}

	return r.Products, nil
}

func (r *ProductRepositoryMemory) Find(id string) (*entities.Product, error) {
	if r.findFailed {
		return nil, errors.NewInternalServerError(errors.FIND_PRODUCT)
	}

	for _, product := range r.Products {
		if product.ID == id {
			return product, nil
		}
	}

	return nil, nil
}

func (r *ProductRepositoryMemory) FailCreate() *ProductRepositoryMemory {
	r.createFailed = true

	return r
}

func (r *ProductRepositoryMemory) FailFindAll() *ProductRepositoryMemory {
	r.findAllFailed = true

	return r
}

func (r *ProductRepositoryMemory) FailFind() *ProductRepositoryMemory {
	r.findFailed = true

	return r
}
