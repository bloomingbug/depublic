package binder

type RegisterRequest struct {
	Name      string `form:"name" validate:"required"`
	Email     string `form:"email" validate:"required,email"`
	Password  string `form:"password" validate:"required,min=8"`
	Phone     string `form:"phone"`
	Avatar    string `form:"avatar"`
	Address   string `form:"address"`
	Birthdate string `form:"birthdate"`
	Gender    string `form:"gender"`
	Token     string `form:"token"`
}

type LoginRequest struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}
