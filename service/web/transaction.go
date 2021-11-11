package web

import (
	"context"
	"errors"
	"fmt"
	"log"

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
	}

	result, err := svc.store.Transaction.Create(ctx, transaction)
	if err != nil {
		return nil, fmt.Errorf("svc.Transaction.Create error: %w", err)
	}

	return result, nil
}

// Balance of current user.
func (svc TransactionWebService) Balance(ctx context.Context, userID int) (*model.BalanceResponse, error) {
	var balanceResponse model.BalanceResponse

	balanceCurrent, err := svc.store.Transaction.Balance(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("svc.Transaction.Balance error: %w", err)
	}
	balanceWidhdraw, err := svc.store.Transaction.BalanceWidhdraw(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("svc.Transaction.Balance error: %w", err)
	}
	fmt.Println(">", balanceWidhdraw)

	balanceResponse.Current = balanceCurrent
	balanceResponse.Withdrawn = balanceWidhdraw

	return &balanceResponse, nil
}

// Withdraw by order transaction.
func (svc TransactionWebService) Withdraw(ctx context.Context, userID int, orderNumber string, amount float64) error {
	balance, err := svc.store.Transaction.Balance(ctx, userID)
	if err != nil {
		return fmt.Errorf("svc.Transaction.Withdraw ифдфтсуerror: %w", err)
	}

	if balance < amount {
		return errors.New("balance amount not enough")
	}

	order := &model.Order{
		UserID: userID,
		Status: model.OrderStatusProcessed,
		Number: orderNumber,
	}

	orderNew, err := svc.store.Order.Create(ctx, order)
	if err != nil {
		return fmt.Errorf("svc.Transaction.Withdraw error: %w", err)
	}

	transaction := &model.Transaction{
		UserID:  userID,
		OrderID: orderNew.ID,
		Amount:  amount,
	}

	_, err = svc.store.Transaction.Create(ctx, transaction)
	if err != nil {
		log.Fatalf("error while create new transaction: %s", err)
	}

	return nil
}
