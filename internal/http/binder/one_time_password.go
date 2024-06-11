package binder

type GenerateOTPRequest struct {
    Email string `json:"email" validate:"required,email"`
}

type FindOTPRequest struct {
    OTPCode string `json:"otp_code" validate:"required,length:8"`
    Email string `json:"email" validate:"required,email"`
}

