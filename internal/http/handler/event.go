package handler

import (
	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/service"
	"github.com/bloomingbug/depublic/pkg/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type EventHandler struct {
	eventService service.EventService
}

func (h *EventHandler) GetAllEvent(c echo.Context) error {
	paginateReq := new(binder.PaginateRequest)
	if err := c.Bind(paginateReq); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	page := 1
	limit := 10

	if paginateReq.Page != nil {
		page = *paginateReq.Page
	}
	if paginateReq.Limit != nil {
		limit = *paginateReq.Limit
	}

	// Create an entity.Paginate object
	paginate := binder.PaginateRequest{
		Page:  &page,
		Limit: &limit,
	}

	events, err := h.eventService.GetAllEventWithPaginate(c, paginate)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK,
		true,
		"sukses menampilkan semua data event",
		events))
}

func NewEventHandler(eventService service.EventService) EventHandler {
	return EventHandler{eventService: eventService}
}
