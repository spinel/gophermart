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
	List(ctx context.Context, userID int) ([]model.Order, error)
}
