package service

import (
	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/repository"
	"github.com/bloomingbug/depublic/internal/util"
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

func (s *eventService) GetAllEventWithPaginateAndFilter(c echo.Context,
	paginate *binder.PaginateRequest,
	filter *binder.FilterRequest,
	sort *binder.SortRequest) (*map[string]interface{}, error) {
	events, totalItems, err := s.eventReposioty.GetAllWithPaginateAndFilter(c.Request().Context(), *paginate, *filter, *sort)
	if err != nil {
		return nil, err
	}

	totalPages := int((totalItems + int64(*paginate.Limit) - 1) / int64(*paginate.Limit))

	data := util.NewPagination(*paginate.Limit, *paginate.Page, int(totalItems), totalPages, events).Response()

	return &data, nil
}

func (s *eventService) FindEventById(c echo.Context, id string) (*entity.Event, error) {
	return s.eventReposioty.FindById(c.Request().Context(), id)
}

type EventService interface {
	GetAllEvent(c echo.Context) (*[]entity.Event, error)
	GetAllEventWithPaginateAndFilter(c echo.Context, paginate *binder.PaginateRequest, filter *binder.FilterRequest, sort *binder.SortRequest) (*map[string]interface{}, error)
	FindEventById(c echo.Context, id string) (*entity.Event, error)
}

func NewEventService(eventRepository repository.EventRepository) EventService {
	return &eventService{
		eventReposioty: eventRepository,
	}
}
