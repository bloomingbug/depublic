package repository

import (
	"context"

	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type oneTimePasswordRepository struct {
	db *gorm.DB
}

func (r *oneTimePasswordRepository) Generate(ctx context.Context, otp *entity.OneTimePassword) (*entity.OneTimePassword, error) {
	if err := r.db.WithContext(ctx).Create(&otp).Error; err != nil {
		return otp, err
	}

	return otp, nil
}

func (r *oneTimePasswordRepository) FindOneByCodeAndEmail(ctx context.Context, email, code string) (*entity.OneTimePassword, error) {
	otp := new(entity.OneTimePassword)
	if err := r.db.WithContext(ctx).Where("email = ? AND otp_code = ? AND is_valid = true", email, code, true).Take(&otp).Error; err != nil {
		return otp, err
	}

	return otp, nil
}

func (r *oneTimePasswordRepository) Used(ctx context.Context, id uuid.UUID) (bool, error) {
	if err := r.db.WithContext(ctx).Model(&entity.OneTimePassword{}).Where("id = ?", id).Update("is_valid", false).Error; err != nil {
		return false, err
	}
	return true, nil
}

type OneTimePasswordRepository interface {
	Generate(ctx context.Context, otp *entity.OneTimePassword) (*entity.OneTimePassword, error)
	FindOneByCodeAndEmail(ctx context.Context, email, code string) (*entity.OneTimePassword, error)
	Used(ctx context.Context, id uuid.UUID) (bool, error)
}

func NewOneTimePasswordRepository(db *gorm.DB) OneTimePasswordRepository {
	return &oneTimePasswordRepository{db}
}
