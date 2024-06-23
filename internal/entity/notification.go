package entity

import "github.com/google/uuid"

type Notification struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Title  string    `json:"title"`
	Detail string    `json:"detail"`
	IsRead bool      `json:"isRead"`
	Auditable
}

func NewNotification(userId uuid.UUID, title, detail string) *Notification {
	return &Notification{
		UserID: userId,
		Title:  title,
		Detail: detail,
		IsRead: false,
	}
}

func ReadNotification(id uuid.UUID) *Notification {
	return &Notification{
		ID:     id,
		IsRead: true,
	}
}
