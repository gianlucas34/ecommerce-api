package factories

import (
	usecases "github.com/gianlucas34/ecommerce-api/internal/application/usecases/user"
	handlers "github.com/gianlucas34/ecommerce-api/internal/infra/api/handlers/user"
	memoryrepos "github.com/gianlucas34/ecommerce-api/internal/infra/repositories/memory"
	"github.com/gianlucas34/ecommerce-api/internal/mocks"
)

func CreateUserHandlerFactory() *handlers.CreateUserHandler {
	userRepository := memoryrepos.NewUserRepositoryMemory()
	passwordHashing := mocks.NewPasswordHashingMock()
	createUserUsecase := usecases.NewCreateUser(userRepository, passwordHashing)
	createUserHandler := handlers.NewCreateUserHandler(createUserUsecase)

	return createUserHandler
}
