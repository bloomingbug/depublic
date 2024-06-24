package handler

import (
	"net/http"

	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/service"
	"github.com/bloomingbug/depublic/pkg/response"
	"github.com/labstack/echo/v4"
)

type LocationHandler struct {
	locationService service.LocationService
}

func (h *LocationHandler) CreateLocation(c echo.Context) error {
	req := new(binder.CreateLocationRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(
			http.StatusBadRequest,
			false,
			"Invalid request data",
		))
	}

	location := &entity.Location{
		Name: req.Name,
	}

	createdLocation, err := h.locationService.CreateLocation(c, location)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error(
			http.StatusInternalServerError,
			false,
			err.Error(),
		))
	}

	return c.JSON(http.StatusOK, response.Success(
		http.StatusOK,
		true,
		"Location created successfully",
		createdLocation,
	))
}

func (h *LocationHandler) GetLocationByID(c echo.Context) error {
	id := c.Param("id")

	location, err := h.locationService.GetLocationByID(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error(
			http.StatusInternalServerError,
			false,
			err.Error(),
		))
	}

	if location == nil {
		return c.JSON(http.StatusNotFound, response.Error(
			http.StatusNotFound,
			false,
			"Location not found",
		))
	}

	return c.JSON(http.StatusOK, response.Success(
		http.StatusOK,
		true,
		"Location retrieved successfully",
		location,
	))
}

func (h *LocationHandler) UpdateLocation(c echo.Context) error {
	req := new(binder.UpdateLocationRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(
			http.StatusBadRequest,
			false,
			"Invalid request data",
		))
	}

	location := &entity.Location{
		ID:   req.ID,
		Name: req.Name,
	}

	updatedLocation, err := h.locationService.UpdateLocation(c, location)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error(
			http.StatusInternalServerError,
			false,
			err.Error(),
		))
	}

	return c.JSON(http.StatusOK, response.Success(
		http.StatusOK,
		true,
		"Location updated successfully",
		updatedLocation,
	))
}

func (h *LocationHandler) DeleteLocation(c echo.Context) error {
	id := c.Param("id")

	err := h.locationService.DeleteLocation(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error(
			http.StatusInternalServerError,
			false,
			err.Error(),
		))
	}

	return c.JSON(http.StatusOK, response.Success(
		http.StatusOK,
		true,
		"Location deleted successfully",
		nil,
	))
}

func NewLocationHandler(locationService service.LocationService) LocationHandler {
	return LocationHandler{locationService}
}
