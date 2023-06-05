package helper

type LoginResponse struct {
	StatusCode int    `json:"-"`
	Error      bool   `json:"error"`
	Message    string `json:"message"`
	Token      string `json:"token,omitempty"`
}

func NewLoginResponse(statusCode int, message string, token string) *LoginResponse {
	return &LoginResponse{
		StatusCode: statusCode,
		Error:      false,
		Message:    message,
		Token:      token,
	}
}
