package repository

import (
	"context"
	"reflect"

	"github.com/bloomingbug/depublic/internal/entity"
	"gorm.io/gorm"
)

type locationRepository struct {
	db *gorm.DB
}

func (r *locationRepository) Create(c context.Context, location *entity.Location) (*entity.Location, error) {
	if err := r.db.WithContext(c).Create(location).Error; err != nil {
		return nil, err
	}
	return location, nil
}

func (r *locationRepository) FindByID(c context.Context, id string) (*entity.Location, error) {
	location := new(entity.Location)
	if err := r.db.WithContext(c).Where("id = ?", id).Take(location).Error; err != nil {
		return nil, err
	}
	return location, nil
}

func (r *locationRepository) Edit(c context.Context, location *entity.Location) (*entity.Location, error) {
	var fields entity.Location

	val := reflect.ValueOf(location).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		if !field.IsZero() {
			reflect.ValueOf(&fields).Elem().FieldByName(fieldName).Set(field)
		}
	}

	if err := r.db.WithContext(c).Model(location).Where("id = ?", location.ID).Updates(fields).Error; err != nil {
		return nil, err
	}
	return location, nil
}

func (r *locationRepository) Delete(c context.Context, location *entity.Location) error {
	if err := r.db.WithContext(c).Delete(location, "id = ?", location.ID).Error; err != nil {
		return err
	}
	return nil
}

type LocationRepository interface {
	Create(c context.Context, location *entity.Location) (*entity.Location, error)
	FindByID(c context.Context, id string) (*entity.Location, error)
	Edit(c context.Context, location *entity.Location) (*entity.Location, error)
	Delete(c context.Context, location *entity.Location) error
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &locationRepository{db: db}
}
