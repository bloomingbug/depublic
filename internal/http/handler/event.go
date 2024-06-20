package handler

import (
	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/service"
	"github.com/bloomingbug/depublic/pkg/response"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type EventHandler struct {
	eventService service.EventService
}

func (h *EventHandler) GetAllEvent(c echo.Context) error {
	paginateReq := new(binder.PaginateRequest)
	if err := c.Bind(paginateReq); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	page := h.getDefaultInt(paginateReq.Page, 1)
	limit := h.getDefaultInt(paginateReq.Limit, 10)
	paginate := &binder.PaginateRequest{
		Page:  &page,
		Limit: &limit,
	}

	filterReq := new(binder.FilterRequest)
	if err := c.Bind(filterReq); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	filter := &binder.FilterRequest{
		Keyword:  filterReq.Keyword,
		Location: filterReq.Location,
		Topic:    filterReq.Topic,
		Category: filterReq.Category,
		Time:     filterReq.Time,
		IsPaid:   filterReq.IsPaid,
	}

	sortReq := new(binder.SortRequest)
	if err := c.Bind(sortReq); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	sort := &binder.SortRequest{
		Sort: sortReq.Sort,
	}

	events, err := h.eventService.GetAllEventWithPaginateAndFilter(c, paginate, filter, sort)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK,
		true,
		"sukses menampilkan semua data event",
		events))
}

func (h *EventHandler) GetDetailEvent(c echo.Context) error {
	id := c.Param("id")
	eventId, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Error(http.StatusNotFound, false, err.Error()))
	}

	event, err := h.eventService.FindEventById(c, eventId)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Error(http.StatusNotFound, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK, true, "sukses menampilkan detail event", event))
}

func (h *EventHandler) getDefaultInt(value *int, defaultValue int) int {
	if value != nil {
		return *value
	}
	return defaultValue
}

func NewEventHandler(eventService service.EventService) EventHandler {
	return EventHandler{eventService: eventService}
}
