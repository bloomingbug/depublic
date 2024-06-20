package service

import (
	"errors"
	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ticketService struct {
	ticketRepo repository.TicketRepository
}

func (s *ticketService) CreateBatchTicket(c echo.Context, transactionID uuid.UUID, tickets *[]entity.Ticket) (*[]entity.Ticket, error) {
	if tickets == nil {
		return nil, errors.New("ticket tidak ada")
	}

	for _, ticket := range *tickets {
		ticket.TransactionID = transactionID
	}

	tickets, err := s.ticketRepo.Creates(c.Request().Context(), tickets)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

type TicketService interface {
	CreateBatchTicket(c echo.Context, transactionID uuid.UUID, tickets *[]entity.Ticket) (*[]entity.Ticket, error)
}

func NewTicketService(ticketRepo repository.TicketRepository) TicketService {
	return &ticketService{ticketRepo: ticketRepo}
}
