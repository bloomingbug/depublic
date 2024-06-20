package repository

import (
	"context"
	"github.com/bloomingbug/depublic/internal/entity"
	"gorm.io/gorm"
	"reflect"
)

type transactionRepository struct {
	db *gorm.DB
}

func (r *transactionRepository) Create(c context.Context, transaction *entity.Transaction) (*entity.Transaction, error) {
	if err := r.db.WithContext(c).Create(&transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *transactionRepository) FindByInvoice(c context.Context, invoice string) (*entity.Transaction, error) {
	transaction := new(entity.Transaction)
	if err := r.db.WithContext(c).Where("invoice = ?", invoice).Take(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *transactionRepository) Edit(c context.Context, transaction *entity.Transaction) (*entity.Transaction, error) {
	var fields entity.Transaction

	val := reflect.ValueOf(transaction).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		if !field.IsZero() {
			reflect.ValueOf(&fields).Elem().FieldByName(fieldName).Set(field)
		}
	}

	if err := r.db.WithContext(c).Model(&transaction).Where("id = ?", transaction.ID).Updates(fields).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

type TransactionRepository interface {
	Create(c context.Context, transaction *entity.Transaction) (*entity.Transaction, error)
	FindByInvoice(c context.Context, invoice string) (*entity.Transaction, error)
	Edit(c context.Context, transaction *entity.Transaction) (*entity.Transaction, error)
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}
