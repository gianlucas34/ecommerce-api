package user

import (
	"encoding/json"
	"net/http"

	"github.com/gianlucas34/ecommerce-api/internal/application/usecases/user"
	"github.com/gianlucas34/ecommerce-api/internal/errors"
)

type CreateUserHandler struct {
	CreateUserUsecase *user.CreateUserUsecase
}

func NewCreateUserHandler(createUserUsecase *user.CreateUserUsecase) *CreateUserHandler {
	return &CreateUserHandler{
		CreateUserUsecase: createUserUsecase,
	}
}

func (h *CreateUserHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var input user.CreateUserUsecaseInput

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
