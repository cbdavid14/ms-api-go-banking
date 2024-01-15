package errs

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (err AppError) AsMessage() *AppError {
	return &AppError{
		Message: err.Message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    404,
		Message: message,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Code:    500,
		Message: message,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Code:    400,
		Message: message,
	}
}
