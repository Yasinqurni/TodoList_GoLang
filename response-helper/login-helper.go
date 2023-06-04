package helper

type LoginResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

func NewLoginResponse(message string, token string) *LoginResponse {
	return &LoginResponse{
		Error:   false,
		Message: message,
		Token:   token,
	}
}
