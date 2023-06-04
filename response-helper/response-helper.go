package helper

type Response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(message string, data interface{}) *Response {
	return &Response{
		Error:   false,
		Message: message,
		Data:    data,
	}
}
