package auth

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginReturnDto struct {
	AccessToken string `json:"access_token"`
}
