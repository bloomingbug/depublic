package repository

import (
	"context"

	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Create(c context.Context, user *entity.User) (*entity.User, error) {
	if err := r.db.WithContext(c).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindByEmail(c context.Context, email string) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.WithContext(c).Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Edit(c context.Context, user *entity.User) (*entity.User, error) {
	if err := r.db.WithContext(c).Model(&entity.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Delete(c context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(c).Delete(&entity.User{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

type UserRepository interface {
	Create(c context.Context, user *entity.User) (*entity.User, error)
	FindByEmail(c context.Context, email string) (*entity.User, error)
	Edit(c context.Context, user *entity.User) (*entity.User, error)
	Delete(c context.Context, id uuid.UUID) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
