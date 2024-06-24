package repository

import (
	"context"

	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRepository interface {
	Create(ctx context.Context, event *entity.Event) (*entity.Event, error)
	FindAll(ctx context.Context) ([]*entity.Event, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Event, error)
	Update(ctx context.Context, event *entity.Event) (*entity.Event, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) Create(ctx context.Context, event *entity.Event) (*entity.Event, error) {
	if err := r.db.WithContext(ctx).Create(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (r *eventRepository) FindAll(ctx context.Context) ([]*entity.Event, error) {
	var events []*entity.Event
	if err := r.db.WithContext(ctx).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *eventRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Event, error) {
	var event entity.Event
	if err := r.db.WithContext(ctx).First(&event, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) Update(ctx context.Context, event *entity.Event) (*entity.Event, error) {
	if err := r.db.WithContext(ctx).Save(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (r *eventRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Event{}).Error; err != nil {
		return err
	}
	return nil
}
