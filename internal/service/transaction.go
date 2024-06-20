package service

import (
	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type transactionService struct {
	transactionRepo repository.TransactionRepository
}

func (s *transactionService) CreateTransaction(c echo.Context, transaction *entity.Transaction) (*entity.Transaction, error) {
	transaction, err := s.transactionRepo.Create(c.Request().Context(), transaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *transactionService) FindTransactionById(c echo.Context, id uuid.UUID) (*entity.Transaction, error) {
	transaction, err := s.transactionRepo.FindById(c.Request().Context(), id)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *transactionService) FindTransactionByInvoice(c echo.Context, invoice string) (*entity.Transaction, error) {
	transaction, err := s.transactionRepo.FindByInvoice(c.Request().Context(), invoice)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *transactionService) EditTransaction(c echo.Context, transaction *entity.Transaction) (*entity.Transaction, error) {
	transaction, err := s.transactionRepo.Edit(c.Request().Context(), transaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

type TransactionService interface {
	CreateTransaction(c echo.Context, transaction *entity.Transaction) (*entity.Transaction, error)
	FindTransactionById(c echo.Context, id uuid.UUID) (*entity.Transaction, error)
	FindTransactionByInvoice(c echo.Context, invoice string) (*entity.Transaction, error)
	EditTransaction(c echo.Context, transaction *entity.Transaction) (*entity.Transaction, error)
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
	}
}
