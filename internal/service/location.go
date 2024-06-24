package service

import (
	"errors"

	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/repository"
	"github.com/labstack/echo/v4"
)

type locationService struct {
	locationRepository repository.LocationRepository
}

func (s *locationService) CreateLocation(c echo.Context, location *entity.Location) (*entity.Location, error) {
	createdLocation, err := s.locationRepository.Create(c.Request().Context(), location)
	if err != nil {
		return nil, err
	}
	return createdLocation, nil
}

func (s *locationService) GetLocationByID(c echo.Context, id string) (*entity.Location, error) {
	location, err := s.locationRepository.FindByID(c.Request().Context(), id)
	if err != nil {
		return nil, err
	}
	return location, nil
}

func (s *locationService) UpdateLocation(c echo.Context, location *entity.Location) (*entity.Location, error) {
	updatedLocation, err := s.locationRepository.Edit(c.Request().Context(), location)
	if err != nil {
		return nil, err
	}
	return updatedLocation, nil
}

func (s *locationService) DeleteLocation(c echo.Context, id string) error {
	location, err := s.locationRepository.FindByID(c.Request().Context(), id)
	if err != nil {
		return err
	}
	if location == nil {
		return errors.New("location not found")
	}
	return s.locationRepository.Delete(c.Request().Context(), location)
}

type LocationService interface {
	CreateLocation(c echo.Context, location *entity.Location) (*entity.Location, error)
	GetLocationByID(c echo.Context, id string) (*entity.Location, error)
	UpdateLocation(c echo.Context, location *entity.Location) (*entity.Location, error)
	DeleteLocation(c echo.Context, id string) error
}

func NewLocationService(locationRepository repository.LocationRepository) LocationService {
	return &locationService{locationRepository: locationRepository}
}
