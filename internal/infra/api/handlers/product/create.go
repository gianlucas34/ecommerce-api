package product

import (
	"encoding/json"
	"net/http"

	"github.com/gianlucas34/ecommerce-api/internal/application/usecases/product"
	"github.com/gianlucas34/ecommerce-api/internal/errors"
)

type CreateProductHandler struct {
	CreateProductUsecase *product.CreateProductUsecase
}

func NewCreateProductHandler(createProductUsecase *product.CreateProductUsecase) *CreateProductHandler {
	return &CreateProductHandler{
		CreateProductUsecase: createProductUsecase,
	}
}

func (h *CreateProductHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var input product.CreateProductUsecaseInput

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		return errors.NewInternalServerError(errors.DECODE)
	}

	err = h.CreateProductUsecase.Execute(input)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)

	return nil
}
