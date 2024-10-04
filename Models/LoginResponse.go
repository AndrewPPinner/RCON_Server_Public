package models

type LoginResponse struct {
	Message string
	Token   string
}

func NewRegisterResponse(message string, token string) LoginResponse {
	return LoginResponse{message, token}
}
