package helper

type ErrorResponse struct {
	StatusCode int    `json:"-"`
	Error      bool   `json:"error"`
	Message    string `json:"message"`
	Err        error  `json:"err,omitempty"`
}

func NewErrorResponse(statusCode int, message string, err error) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: statusCode,
		Error:      true,
		Message:    message,
		Err:        err,
	}
}
