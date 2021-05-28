package user

type addNameRequest struct {
	Email    string `json:"email"    validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name`
}

type jwtRequest struct {
	Email    string `json:"email"    validate:"required"`
	Password string `json:"password" validate:"required"`
}

type restrictedRequest struct {
	Token string `json:"token" validate:"required"`
}