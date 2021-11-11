package store

import (
	"context"

	"github.com/spinel/gophermart/model"
)

// UserRepo is a store for Users.
type UserRepo interface {
	Create(ctx context.Context, reqUser *model.User) (*model.User, error)
	GetByLogin(ctx context.Context, login string) (*model.User, error)
}

// OrderRepo is a store for orders.
type OrderRepo interface {
	Create(ctx context.Context, order *model.Order) (*model.Order, error)
	Update(ctx context.Context, order *model.Order) error
	List(ctx context.Context, userID int) ([]model.Order, error)
	GetByStatus(ctx context.Context, orderStatus string) ([]model.Order, error)
	GetByNumber(ctx context.Context, orderNumber string) (*model.Order, error)
}

// TransactionRepo is a store for transactions.
type TransactionRepo interface {
	Create(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error)
	Balance(ctx context.Context, userID int) (float64, error)
	BalanceWidhdraw(ctx context.Context, userID int) (float64, error)
}
