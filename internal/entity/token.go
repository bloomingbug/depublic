package entity

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	ID        uuid.UUID `json:"id"`
	Token     string    `json:"token"`
	Email     string    `json:"email"`
	Action    string    `json:"action"`
	ExpiresAt time.Time `json:"expires_at" sql:"expires_at"`
	Auditable
}

func NewTokenRegister(token, email string) *Token {
	return &Token{
		ID:        uuid.New(),
		Token:     token,
		Email:     email,
		Action:    "register",
		ExpiresAt: time.Now().Add(time.Minute * 15),
		Auditable: NewAuditable(),
	}
}

func NewTokenForgotPassword(token, email string) *Token {
	return &Token{
		ID:        uuid.New(),
		Token:     token,
		Email:     email,
		Action:    "forgot-password",
		ExpiresAt: time.Now().Add(time.Minute * 15),
		Auditable: NewAuditable(),
	}
}
