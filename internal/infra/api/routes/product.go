package routes

import (
	"net/http"

	"github.com/gianlucas34/ecommerce-api/internal/infra/api/handlers/product"
	"github.com/gianlucas34/ecommerce-api/internal/infra/api/middlewares"
)

type ProductRoutes struct {
	Mux                  *http.ServeMux
	CreateProductHandler *product.CreateProductHandler
}

func NewProductRoutes(mux *http.ServeMux, createProductHandler *product.CreateProductHandler) *ProductRoutes {
	return &ProductRoutes{
		Mux:                  mux,
		CreateProductHandler: createProductHandler,
	}
}

func (r *ProductRoutes) Register() {
	r.Mux.HandleFunc("POST /products", middlewares.ErrorHandlerMiddleware(r.CreateProductHandler.Handle))
}
