package binder

type VerifyOTPRequest struct {
	Email   string `json:"email" query:"email"`
	OTPCode string `json:"otp_code" query:"otp_code"`
}
