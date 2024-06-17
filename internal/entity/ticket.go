package entity

import (
	"github.com/google/uuid"
	"time"
)

type Ticket struct {
	ID            uuid.UUID   `json:"id"`
	NoTicket      string      `json:"no_ticket"`
	Name          string      `json:"name"`
	PersonalNo    *string     `json:"personal_no,omitempty"`
	Birthdate     time.Time   `json:"birthdate"`
	Phone         *string     `json:"phone,omitempty"`
	Email         *string     `json:"email"`
	Gender        *Gender     `json:"gender,omitempty"`
	Price         int64       `json:"price"`
	TransactionID uuid.UUID   `json:"-"`
	Transaction   Transaction `json:"transaction"`
	TimetableID   uuid.UUID   `json:"-"`
	Timetable     Timetable   `json:"timetable"`
	Auditable
}

func NewTicket(noTicket, name string,
	personalNo *string,
	birthdate time.Time, phone,
	email *string,
	gender *Gender,
	price int64) *Ticket {
	return &Ticket{
		ID:         uuid.New(),
		NoTicket:   noTicket,
		Name:       name,
		PersonalNo: personalNo,
		Birthdate:  birthdate,
		Phone:      phone,
		Email:      email,
		Gender:     gender,
		Price:      price,
	}
}
