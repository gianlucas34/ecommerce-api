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

	repositories := factories.NewRepositoriesFactory()
	adapters := factories.NewAdaptersFactory()
	usecases := factories.NewUsecasesFactory(repositories, adapters)
	handlers := factories.NewHandlersFactory(usecases)

	userRoutes := routes.NewUserRoutes(router, handlers.CreateUserHandler)
	productRoutes := routes.NewProductRoutes(router, handlers.CreateProductHandler)

	userRoutes.Register()
	productRoutes.Register()

	return http.ListenAndServe(":8080", router)
}
