package model

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
