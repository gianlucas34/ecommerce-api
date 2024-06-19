package userhdls

import (
	"encoding/json"
	"net/http"

	userucs "github.com/gianlucas34/ecommerce-api/internal/application/usecases/user"
	"github.com/gianlucas34/ecommerce-api/internal/errors"
)

type CreateUserHandler struct {
	CreateUserUsecase *userucs.CreateUser
}

func NewCreateUserHandler(createUserUsecase *userucs.CreateUser) *CreateUserHandler {
	return &CreateUserHandler{
		CreateUserUsecase: createUserUsecase,
	}
}

func (h *CreateUserHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var input userucs.CreateUserInput

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		return errors.NewInternalServerError(errors.DECODE)
	}

	err = h.CreateUserUsecase.Execute(input)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)

	return nil
}
