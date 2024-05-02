package errors

const CREATE_USER = "Erro ao criar usuário!"
const FIND_ALL_USERS = "Erro ao buscar usuários!"
const FIND_USER = "Erro ao buscar usuário!"
const FIND_USER_BY_EMAIL = "Erro ao buscar usuário por e-mail!"

const GENERATE_HASH = "Erro ao gerar o hash da senha!"
const COMPARE_PASSWORD = "Erro ao comparar a senha com o hash!"

type InternalServerError struct {
	Message string
	Code    int
}

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{
		Message: message,
		Code:    500,
	}
}

func (e *InternalServerError) Error() string {
	return e.Message
}
