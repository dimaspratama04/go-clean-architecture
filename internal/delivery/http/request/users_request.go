package request

type LoginRequest struct {
	Username string `json:"username" validate:"required" require:"true"`
	Email    string `json:"email" validate:"required" require:"true"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required" require:"true"`
	Email    string `json:"email" validate:"required" require:"true"`
}
