package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/gianlucas34/ecommerce-api/internal/errors"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

type ErrorResponse struct {
	Error string `json:"error"`
}

func ErrorHandlerMiddleware(handler HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)

		if err != nil {
			switch err.(type) {
			case *errors.ValidationError:
				w.WriteHeader(http.StatusBadRequest)
			case *errors.FieldInUseError:
				w.WriteHeader(http.StatusUnprocessableEntity)
			case *errors.InternalServerError:
				w.WriteHeader(http.StatusInternalServerError)
			}

			json, _ := json.Marshal(ErrorResponse{
				Error: err.Error(),
			})

			w.Write(json)
		}
	})
}
