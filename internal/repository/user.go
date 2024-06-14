package repository

import (
	"context"

	"github.com/bloomingbug/depublic/internal/entity"
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
	fields := make(map[string]interface{})

	if user.Email != "" {
		fields["email"] = user.Email
	}
	if user.Password != "" {
		fields["password"] = user.Password
	}
	if user.Role != "" {
		fields["role"] = user.Role
	}
	if user.Avatar != "" {
		fields["avatar"] = user.Avatar
	}
	if user.Address != "" {
		fields["address"] = user.Address
	}
	if user.Birthdate != nil {
		fields["birthdate"] = user.Birthdate
	}
	if user.Gender != "" {
		fields["gender"] = user.Gender
	}
	if user.Phone != "" {
		fields["phone"] = user.Phone
	}
	if user.Name != "" {
		fields["name"] = user.Name
	}

	if err := r.db.WithContext(c).Model(&user).Where("id = ?", user.ID).Updates(fields).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Delete(c context.Context, user *entity.User) error {
	if err := r.db.WithContext(c).Delete(&user, "id = ?", user.ID).Error; err != nil {
		return err
	}
	return nil
}

type UserRepository interface {
	Create(c context.Context, user *entity.User) (*entity.User, error)
	FindByEmail(c context.Context, email string) (*entity.User, error)
	Edit(c context.Context, user *entity.User) (*entity.User, error)
	Delete(c context.Context, user *entity.User) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
