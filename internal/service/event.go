package service

import (
	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/repository"
	"github.com/labstack/echo/v4"
)

type eventService struct {
	eventReposioty repository.EventRepository
}

func (s *eventService) GetAllEvent(c echo.Context) (*[]entity.Event, error) {
	events, err := s.eventReposioty.GetAll(c.Request().Context())
	if err != nil {
		return nil, err
	}

	return &events, nil
}

func (s *eventService) GetAllEventWithPaginate(c echo.Context, paginate binder.PaginateRequest) (*map[string]interface{}, error) {
	events, err := s.eventReposioty.GetAllPaginate(c.Request().Context(), *paginate.Page, *paginate.Limit)
	if err != nil {
		return nil, err
	}

	return &events, nil
}

type EventService interface {
	GetAllEvent(c echo.Context) (*[]entity.Event, error)
	GetAllEventWithPaginate(c echo.Context, paginate binder.PaginateRequest) (*map[string]interface{}, error)
}

func NewEventService(eventRepository repository.EventRepository) EventService {
	return &eventService{
		eventReposioty: eventRepository,
	}
}
