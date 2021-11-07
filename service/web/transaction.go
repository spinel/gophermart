package web

import (
	"context"
	"fmt"

	"github.com/spinel/gophermart/model"
	"github.com/spinel/gophermart/store"
)

// TransactionWebService ...
type TransactionWebService struct {
	ctx   context.Context
	store *store.Store
}

// NewTransactionWebService is a transaction service.
func NewTransactionWebService(ctx context.Context, store *store.Store) *TransactionWebService {
	return &TransactionWebService{
		ctx:   ctx,
		store: store,
	}
}

// Create transaction service.
func (svc TransactionWebService) Create(ctx context.Context, userID int, orderNumber int) (*model.Transaction, error) {
	transaction := &model.Transaction{
		OrderID: 1,
		UserID:  userID,
		Status:  model.OrderStatusNew,
		//Number:  orderNumber,
	}

	result, err := svc.store.Transaction.Create(ctx, transaction)
	if err != nil {
		return nil, fmt.Errorf("svc.Transaction.Create error: %w", err)
	}

	return result, nil
}

// Balance of current user.
func (svc TransactionWebService) Balance(ctx context.Context, userID int) (float64, error) {
	balance, err := svc.store.Transaction.Balance(ctx, userID)
	if err != nil {
		return 0, fmt.Errorf("svc.Transaction.Balance error: %w", err)
	}

	return balance, nil
}

// Withdraw by order transaction.
func (svc TransactionWebService) Withdraw(ctx context.Context, userID int, order string, amount float64) error {
	balance, err := svc.store.Transaction.Balance(ctx, userID)
	if err != nil {
		return fmt.Errorf("svc.Transaction.Balance error: %w", err)
	}

	if balance < amount {
		return fmt.Errorf("svc.Transaction.Balance error: %w", err)
	}

	_ = &model.Transaction{
		UserID:  userID,
		OrderID: 1,
		Amount:  amount,
	}

	return nil
}
