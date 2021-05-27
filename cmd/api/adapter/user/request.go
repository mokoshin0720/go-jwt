package user

type addNameRequest struct {
	ID       int64  `json:"id"       validate:"required"`
	Name     string `json:"name"     validate:"required"`
	Email    string `json:"email"    validate:"required"`
	Password string `json:"password" validate:"required"`
}

type jwtRequest struct {
	Email    string `json:"email"    validate:"required"`
	Password string `json:"password" validate:"required"`
}

type restrictedRequest struct {
	Token string `json:"token" validate:"required"`
}