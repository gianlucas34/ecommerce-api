package errors

const ALREADY_IN_USE_EMAIL = "O e-mail fornecido já está em uso!"

type FieldInUseError struct {
	Message string
	Code    int
}

func NewFieldInUseError(message string) *FieldInUseError {
	return &FieldInUseError{
		Message: message,
		Code:    422,
	}
}

func (e *FieldInUseError) Error() string {
	return e.Message
}
