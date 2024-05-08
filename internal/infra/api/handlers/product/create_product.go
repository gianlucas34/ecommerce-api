package handlers

import (
	"net/http"

	usecases "github.com/gianlucas34/ecommerce-api/internal/application/usecases/product"
)

type CreateProductHandler struct {
	CreateProductUsecase *usecases.CreateProduct
}

func NewCreateProductHandler(createProductUsecase *usecases.CreateProduct) *CreateProductHandler {
	return &CreateProductHandler{
		CreateProductUsecase: createProductUsecase,
	}
}

func (h *CreateProductHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}
