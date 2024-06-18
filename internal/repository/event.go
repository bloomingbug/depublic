package repository

import (
	"context"
	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/util"
	"gorm.io/gorm"
)

type eventRepository struct {
	db *gorm.DB
}

func (r *eventRepository) GetAll(c context.Context) ([]entity.Event, error) {
	events := make([]entity.Event, 0)
	err := r.db.WithContext(c).
		Where("is_public = ? AND is_approved = ?", true, true).Find(&events).Error

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r *eventRepository) GetAllPaginate(c context.Context, page, limit int) (map[string]interface{}, error) {
	var totalItems int64
	events := make([]entity.Event, 0)

	err := r.db.WithContext(c).
		Model(&entity.Event{}).
		Where("is_public = ? AND is_approved = ?", true, true).
		Count(&totalItems).Error
	if err != nil {
		return nil, err
	}

	totalPages := int((totalItems + int64(limit) - 1) / int64(limit))

	if int(totalItems) <= 0 {
		pagination := util.NewPagination(limit, page, int(totalItems), totalPages, []string{}).Response()

		return pagination, nil
	}

	err = r.db.WithContext(c).
		Scopes(util.Paginate(page, limit)).
		Where("is_public = ? AND is_approved = ?", true, true).Find(&events).Error

	if err != nil {
		return nil, err
	}

	pagination := util.NewPagination(limit, page, int(totalItems), totalPages, events).Response()

	return pagination, nil
}

type EventRepository interface {
	GetAll(c context.Context) ([]entity.Event, error)
	GetAllPaginate(c context.Context, page, limit int) (map[string]interface{}, error)
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db: db}
}
