package helper

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Err     error  `json:"err,omitempty"`
}

func NewErrorResponse(message string, err error) *ErrorResponse {
	return &ErrorResponse{
		Error:   true,
		Message: message,
		Err:     err,
	}
}
