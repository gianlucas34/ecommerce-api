package routes

import (
	"net/http"

	handlers "github.com/gianlucas34/ecommerce-api/internal/infra/api/handlers/user"
	"github.com/gianlucas34/ecommerce-api/internal/infra/api/middlewares"
)

type UserRoutes struct {
	Router            *http.ServeMux
	CreateUserHandler *handlers.CreateUserHandler
}

func NewUserRoutes(router *http.ServeMux, createUserHandler *handlers.CreateUserHandler) *UserRoutes {
	return &UserRoutes{
		Router:            router,
		CreateUserHandler: createUserHandler,
	}
}

func (r *UserRoutes) Register() {
	r.Router.HandleFunc("POST /users", middlewares.ErrorHandlerMiddleware(r.CreateUserHandler.Handle))
}
