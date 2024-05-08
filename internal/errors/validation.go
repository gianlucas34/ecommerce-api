package errors

const EMPTY_USERNAME = "O nome de usuário não pode ficar em branco!"
const EMPTY_USER_EMAIL = "O e-mail não pode ficar em branco!"
const INVALID_USER_EMAIL = "O e-mail fornecido não é um e-mail válido!"
const EMPTY_USER_PASSWORD = "A senha não pode ficar em branco!"

const EMPTY_PRODUCT_NAME = "O nome do produto não pode ficar em branco!"
const INVALID_PRODUCT_PRICE = "O preço do produto é inválido, forneça um preço maior que 0!"

type ValidationError struct {
	Message string
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		Message: message,
	}
}

func (e *ValidationError) Error() string {
	return e.Message
}
