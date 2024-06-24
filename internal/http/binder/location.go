package binder

type CreateLocationRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateLocationRequest struct {
	ID   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
