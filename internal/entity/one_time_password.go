package entity

import (
	"time"

	"github.com/google/uuid"
)

type OneTimePassword struct {
	ID        uuid.UUID `json:"id"`
	OTPCode   string    `json:"otp_code" length:"8"`
	Email     string    `json:"email"`
	IsValid   bool      `json:"is_valid"`
	ExpiresAt time.Time `json:"expires_at" sql:"expires_at"`
	Auditable
}

func NewOneTimePassword(otpCode, email string) *OneTimePassword {
	return &OneTimePassword{
		ID:        uuid.New(),
		OTPCode:   otpCode,
		Email:     email,
		IsValid:   true,
		ExpiresAt: time.Now().Add(time.Minute * 15),
		Auditable: NewAuditable(),
	}
}

func UsedOneTimePassword(id uuid.UUID) *OneTimePassword {
	return &OneTimePassword{
		ID:        id,
		IsValid:   false,
		Auditable: UpdateAuditable(),
	}
}
