package api

import (
	"net/http"

	"github.com/gianlucas34/ecommerce-api/internal/infra/api/routes"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	router := http.NewServeMux()

	routes.UserRoutes(router)
	routes.ProductRoutes(router)

	return http.ListenAndServe(":8080", router)
}
