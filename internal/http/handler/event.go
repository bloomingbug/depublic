package handler

import (
	"net/http"

	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	eventService service.EventService
}

func NewEventHandler(eventService service.EventService) *EventHandler {
	return &EventHandler{eventService: eventService}
}

func (h *EventHandler) CreateEvent(c echo.Context) error {
	params := entity.NewEventParams{}
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	createdEvent, err := h.eventService.CreateEvent(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create event")
	}

	return c.JSON(http.StatusCreated, createdEvent)
}

func (h *EventHandler) GetAllEvents(c echo.Context) error {
	events, err := h.eventService.GetAllEvents(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch events")
	}

	return c.JSON(http.StatusOK, events)
}

func (h *EventHandler) GetEventByID(c echo.Context) error {
	eventID := c.Param("id")
	id, err := uuid.Parse(eventID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid event ID")
	}

	event, err := h.eventService.GetEventByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Event not found")
	}

	return c.JSON(http.StatusOK, event)
}

func (h *EventHandler) UpdateEvent(c echo.Context) error {
	params := entity.EditEventParams{}
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	updatedEvent, err := h.eventService.UpdateEvent(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update event")
	}

	return c.JSON(http.StatusOK, updatedEvent)
}

func (h *EventHandler) DeleteEvent(c echo.Context) error {
	eventID := c.Param("id")
	id, err := uuid.Parse(eventID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid event ID")
	}

	if err := h.eventService.DeleteEvent(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete event")
	}

	return c.JSON(http.StatusOK, "Event deleted successfully")
}

func (h *EventHandler) ListEvents(c echo.Context) error {
	events, err := h.eventService.GetAllEvents(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch events")
	}
	return c.JSON(http.StatusOK, events)
}

func (h *EventHandler) ApproveEvent(c echo.Context) error {
	id := c.Param("id")
	eventID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid event ID"})
	}

	event, err := h.eventService.ApproveEvent(c.Request().Context(), eventID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, event)
}

func (h *EventHandler) RejectEvent(c echo.Context) error {
	id := c.Param("id")
	eventID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid event ID"})
	}

	event, err := h.eventService.RejectEvent(c.Request().Context(), eventID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, event)
}
