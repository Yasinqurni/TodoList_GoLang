package helper

type Response struct {
	StatusCode int         `json:"-"`
	Error      bool        `json:"error"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func NewResponse(statusCode int, message string, data interface{}) *Response {
	return &Response{
		StatusCode: statusCode,
		Error:      false,
		Message:    message,
		Data:       data,
	}
}
