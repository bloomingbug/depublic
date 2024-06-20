package repository

import (
	"context"
	"github.com/bloomingbug/depublic/internal/entity"
	"gorm.io/gorm"
)

type ticketRepository struct {
	db *gorm.DB
}

func (r *ticketRepository) Creates(c context.Context, tickets *[]entity.Ticket) (*[]entity.Ticket, error) {
	if err := r.db.WithContext(c).Create(tickets).Error; err != nil {
		return nil, err
	}

	return tickets, nil
}

type TicketRepository interface {
	Creates(c context.Context, tickets *[]entity.Ticket) (*[]entity.Ticket, error)
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db: db}
}
