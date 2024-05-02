package routes

import (
	"net/http"

	handlers "github.com/gianlucas34/ecommerce-api/internal/infra/api/handlers/product"
)

func ProductRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /products", handlers.CreateProductHandler)
}
