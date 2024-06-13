package binder

type RegisterRequest struct {
	Name                 string `form:"name" json:"name" validate:"required,alphanum"`
	Email                string `form:"email" json:"email" validate:"required,email"`
	Password             string `form:"password" json:"password" validate:"required,min=8"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" validate:"required,min=8,eqfield=Password"`
	Phone                string `form:"phone" json:"phone"`
	Avatar               string `form:"avatar" json:"avatar"`
	Address              string `form:"address" json:"address"`
	Birthdate            string `form:"birthdate" json:"birthdate"`
	Gender               string `form:"gender" json:"gender" validate:"required,oneof=M F"`
	Token                string `form:"token" json:"token" validate:"required,uuid"`
}

type LoginRequest struct {
	Email    string `form:"email" json:"email" validate:"required,email"`
	Password string `form:"password" json:"password" validate:"required"`
}
