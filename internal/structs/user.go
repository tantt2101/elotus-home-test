package structs

type RegisterRequest struct {
	Username        string `json:"username" validate:"required,min=12"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type UserResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}