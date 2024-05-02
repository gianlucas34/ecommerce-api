package routes

import (
	"net/http"

	handlers "github.com/gianlucas34/ecommerce-api/internal/infra/api/handlers/user"
)

func UserRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /users", handlers.CreateUserHandler)
}
