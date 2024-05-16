package producthdls

import (
	"encoding/json"
	"net/http"

	usecases "github.com/gianlucas34/ecommerce-api/internal/application/usecases/product"
	"github.com/gianlucas34/ecommerce-api/internal/errors"
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
	var input usecases.CreateProductInput

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
