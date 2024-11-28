package memory

import (
	"github.com/gianlucas34/ecommerce-api/internal/domain"
	"github.com/gianlucas34/ecommerce-api/internal/errors"
)

type ProductRepository struct {
	Products      []*domain.Product
	createFailed  bool
	findAllFailed bool
	findFailed    bool
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		Products: make([]*domain.Product, 0),
	}
}

func (r *ProductRepository) Create(product *domain.Product) error {
	if r.createFailed {
		return errors.NewInternalServerError(errors.CREATE_PRODUCT)
	}

	r.Products = append(r.Products, product)

	return nil
}

func (r *ProductRepository) FindAll() ([]*domain.Product, error) {
	if r.findAllFailed {
		return nil, errors.NewInternalServerError(errors.FIND_ALL_PRODUCTS)
	}

	return r.Products, nil
}

func (r *ProductRepository) Find(id string) (*domain.Product, error) {
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

func (r *ProductRepository) FailCreate() *ProductRepository {
	r.createFailed = true

	return r
}

func (r *ProductRepository) FailFindAll() *ProductRepository {
	r.findAllFailed = true

	return r
}

func (r *ProductRepository) FailFind() *ProductRepository {
	r.findFailed = true

	return r
}
