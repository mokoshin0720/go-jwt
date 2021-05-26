package user

type addNameRequest struct {
	ID   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}