package entity

import "github.com/google/uuid"

type Transaction struct {
	ID         uuid.UUID `json:"id"`
	Invoice    string    `json:"invoice"`
	GrandTotal int64     `json:"grand_total"`
	SnapToken  string    `json:"snap_token"`
	Status     string    `json:"status"`
	UserID     uuid.UUID `json:"-"`
	User       User      `json:"user"`
	Auditable
}
