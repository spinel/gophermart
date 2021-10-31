package web

import (
	"context"
	"fmt"
	"time"

	"github.com/spinel/gophermart/model"
	"github.com/spinel/gophermart/pkg/luhn"
	"github.com/spinel/gophermart/store"
)

// OrderWebService ...
type OrderWebService struct {
	ctx   context.Context
	store *store.Store
}

// NewOrderWebService is an order service.
func NewOrderWebService(ctx context.Context, store *store.Store) *OrderWebService {
	return &OrderWebService{
		ctx:   ctx,
		store: store,
	}
}

// Create order service.
func (svc OrderWebService) Create(ctx context.Context, userID int, orderNumber int) (*model.Order, error) {
	// order number Luhn validation
	if !luhn.Valid(orderNumber) {
		return nil, fmt.Errorf("svc.Order.Create error: Luhn validation failed")
	}

	order := &model.Order{
		UserID:     userID,
		Status:     model.OrderStatusNew,
		Number:     orderNumber,
		UploadedAt: time.Now().UTC(),
	}

	result, err := svc.store.Order.Create(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("svc.Order.Create error: %w", err)
	}

	return result, nil
}

// List of orders of current user.
func (svc OrderWebService) List(ctx context.Context, userID int) ([]model.Order, error) {
	orders, err := svc.store.Order.List(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("svc.Order.Create error: %w", err)
	}

	return orders, nil
}
