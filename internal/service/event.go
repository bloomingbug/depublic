package service

import (
	"context"
	"time"

	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/repository"
	"github.com/google/uuid"
)

type EventService interface {
	CreateEvent(ctx context.Context, params entity.NewEventParams) (*entity.Event, error)
	GetAllEvents(ctx context.Context) ([]*entity.Event, error)
	GetEventByID(ctx context.Context, id uuid.UUID) (*entity.Event, error)
	UpdateEvent(ctx context.Context, params entity.EditEventParams) (*entity.Event, error)
	DeleteEvent(ctx context.Context, id uuid.UUID) error
	ApproveEvent(ctx context.Context, eventID uuid.UUID) (*entity.Event, error)
	RejectEvent(ctx context.Context, eventID uuid.UUID) (*entity.Event, error)
}

type eventService struct {
	eventRepo repository.EventRepository
}

func NewEventService(eventRepo repository.EventRepository) EventService {
	return &eventService{
		eventRepo: eventRepo,
	}
}

func (s *eventService) CreateEvent(ctx context.Context, params entity.NewEventParams) (*entity.Event, error) {
	event := entity.NewEvent(params)
	return s.eventRepo.Create(ctx, event)
}

func (s *eventService) GetAllEvents(ctx context.Context) ([]*entity.Event, error) {
	return s.eventRepo.FindAll(ctx)
}

func (s *eventService) GetEventByID(ctx context.Context, id uuid.UUID) (*entity.Event, error) {
	return s.eventRepo.FindByID(ctx, id)
}

func (s *eventService) UpdateEvent(ctx context.Context, params entity.EditEventParams) (*entity.Event, error) {
	event := entity.EditEvent(params)
	return s.eventRepo.Update(ctx, event)
}

func (s *eventService) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	return s.eventRepo.Delete(ctx, id)
}

func (s *eventService) ApproveEvent(ctx context.Context, eventID uuid.UUID) (*entity.Event, error) {
	event, err := s.eventRepo.FindByID(ctx, eventID)
	if err != nil {
		return nil, err
	}

	// Perform any business logic checks before approving
	event.IsApproved = true
	now := time.Now()
	event.ApprovedAt = &now

	// Save the updated event to the repository
	updatedEvent, err := s.eventRepo.Update(ctx, event)
	if err != nil {
		return nil, err
	}

	return updatedEvent, nil
}

func (s *eventService) RejectEvent(ctx context.Context, eventID uuid.UUID) (*entity.Event, error) {
	event, err := s.eventRepo.FindByID(ctx, eventID)
	if err != nil {
		return nil, err
	}

	// Perform any business logic checks before rejecting
	event.IsApproved = false
	event.ApprovedAt = nil

	// Save the updated event to the repository
	updatedEvent, err := s.eventRepo.Update(ctx, event)
	if err != nil {
		return nil, err
	}

	return updatedEvent, nil
}
