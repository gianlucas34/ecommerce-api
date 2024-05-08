package errors

const DECODE = "Erro ao decodificar dados!"

const CREATE_USER = "Erro ao criar usuário!"
const FIND_ALL_USERS = "Erro ao buscar usuários!"
const FIND_USER = "Erro ao buscar usuário!"
const FIND_USER_BY_EMAIL = "Erro ao buscar usuário por e-mail!"

const CREATE_PRODUCT = "Erro ao criar produto!"
const FIND_ALL_PRODUCTS = "Erro ao buscar produtos!"
const FIND_PRODUCT = "Erro ao buscar produto!"

const GENERATE_HASH = "Erro ao gerar o hash da senha!"
const COMPARE_PASSWORD = "Erro ao comparar a senha com o hash!"

type InternalServerError struct {
	Message string
}

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{
		Message: message,
	}
}

func (e *InternalServerError) Error() string {
	return e.Message
}
