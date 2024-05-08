package errors

const ALREADY_IN_USE_EMAIL = "O e-mail fornecido já está em uso!"

type FieldInUseError struct {
	Message string
}

func NewFieldInUseError(message string) *FieldInUseError {
	return &FieldInUseError{
		Message: message,
	}
}

func (e *FieldInUseError) Error() string {
	return e.Message
}
