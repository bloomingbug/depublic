package service

import (
	"errors"
	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/repository"
	"github.com/bloomingbug/depublic/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type notificationService struct {
	notifRepository repository.NotificationRepository
}

func (s *notificationService) GetUserNotification(c echo.Context, id uuid.UUID, paginate *binder.PaginateRequest, isRead *bool) (*map[string]interface{}, error) {
	notifications, total, err := s.notifRepository.GetByUserIdWithPagination(c.Request().Context(), id, *paginate, isRead)
	if err != nil {
		return nil, err
	}

	totalPages := int((total + int64(*paginate.Limit) - 1) / int64(*paginate.Limit))

	data := util.NewPagination(*paginate.Limit, *paginate.Page, int(total), totalPages, notifications).Response()

	return &data, nil
}

func (s *notificationService) GetDetailNotification(c echo.Context, id uuid.UUID) (*entity.Notification, error) {
	// Check request user from valid user (auth user == user_id)
	notificationOld, err := s.notifRepository.FindById(c.Request().Context(), id)
	if err != nil {
		return nil, err
	}

	if notificationOld.UserID != id {
		return nil, errors.New("tidak memiliki hak untuk mengakses notifikasi ini")
	}

	//	Edit notif is_read to true
	notifDTO := entity.ReadNotification(id)
	err = s.notifRepository.Edit(c.Request().Context(), notifDTO)
	if err != nil {
		return nil, err
	}

	notification, err := s.notifRepository.FindById(c.Request().Context(), id)
	if err != nil {
		return nil, err
	}

	return notification, nil
}

type NotificationService interface {
	GetUserNotification(c echo.Context, id uuid.UUID, paginate *binder.PaginateRequest, isRead *bool) (*map[string]interface{}, error)
	GetDetailNotification(c echo.Context, id uuid.UUID) (*entity.Notification, error)
}

func NewNotificationService(notifRepository repository.NotificationRepository) NotificationService {
	return &notificationService{notifRepository: notifRepository}
}
