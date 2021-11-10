package service

import (
	"context"

	"github.com/spinel/gophermart/model"
)

type UserService interface {
	Create(ctx context.Context, userRegisterForm model.UserRegisterForm) (*model.User, error)
	Login(context.Context, model.UserRegisterForm) (*model.User, error)
}

type OrderService interface {
	Create(ctx context.Context, userID int, orderNumber int) (*model.Order, error)
	List(ctx context.Context, userID int) ([]model.Order, error)
}

type TransactionService interface {
	Create(ctx context.Context, userID int, orderNumber int) (*model.Transaction, error)
	Balance(ctx context.Context, userID int) (float64, error)
	Withdraw(ctx context.Context, userID int, order string, amount float64) error
}

type CacheService interface {
	Do(ctx context.Context, commandName string, args ...interface{}) (reply interface{}, err error)
}

type MemoryService interface {
	Add(ctx context.Context, sessionToken string, userID int)
	Get(ctx context.Context, sessionToken string) int
}
