package api

import (
	"net/http"

	"github.com/gianlucas34/ecommerce-api/internal/infra/api/routes"
	"github.com/gianlucas34/ecommerce-api/internal/infra/factories"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	router := http.NewServeMux()

	createUserHandlerFactory := factories.CreateUserHandlerFactory()
	createProductHandlerFactory := factories.CreateProductHandlerFactory()

	userRoutes := routes.NewUserRoutes(router, createUserHandlerFactory)
	productRoutes := routes.NewProductRoutes(router, createProductHandlerFactory)

	userRoutes.Register()
	productRoutes.Register()

	return http.ListenAndServe(":8080", router)
}
