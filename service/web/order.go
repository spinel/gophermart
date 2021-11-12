package web

import (
	"context"

	"fmt"
	"log"
	"time"

	"github.com/spinel/gophermart/model"
	"github.com/spinel/gophermart/pkg/luhn"
	"github.com/spinel/gophermart/pkg/types"
	"github.com/spinel/gophermart/service/ext"
	"github.com/spinel/gophermart/store"

	"github.com/pkg/errors"
)

// OrderWebService ...
type OrderWebService struct {
	ctx   context.Context
	store *store.Store
	ext   *ext.ExtAccuralService
}

// NewOrderWebService is an order service.
func NewOrderWebService(ctx context.Context, store *store.Store, extService *ext.ExtAccuralService) *OrderWebService {
	orderService := &OrderWebService{
		ctx:   ctx,
		store: store,
		ext:   extService,
	}
	go orderService.workerUpdateStatus(500)

	return orderService

}

func (svc OrderWebService) workerUpdateStatus(interval int) {
	ctx := context.Background()
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	done := make(chan bool)

	for {
		select {
		case <-done:
			fmt.Println("exit!")
			return
		case t := <-ticker.C:

			newOrders, err := svc.store.Order.GetByStatus(ctx, model.OrderStatusNew)
			if err != nil {
				log.Fatalf("error while load new orders: %s", err)
			}

			for _, order := range newOrders {
				fmt.Printf("get %s %s\n", t, order.Number)
				orderResp, err := svc.ext.OrderStatus(ctx, order.Number)
				if err != nil {
					log.Fatalf("external api error: %s", err)
					continue
				}

				if orderResp.Order != "" {
					// set updated data
					order.Accural = orderResp.Accural
					order.Status = orderResp.Status

					// update order db
					err := svc.store.Order.Update(ctx, &order)
					if err != nil {
						log.Fatalf("error while update order(%s): %s", order.Number, err)
					}

					transaction := &model.Transaction{
						OrderID: order.ID,
						UserID:  order.UserID,
						Amount:  order.Accural,
						Type:    model.TransactionTypeRefill,
					}

					_, err = svc.store.Transaction.Create(ctx, transaction)
					if err != nil {
						log.Fatalf("error while create new transaction: %s", err)
					}
				}
			}
		}
	}

}

// Create order service.
func (svc OrderWebService) Create(ctx context.Context, userID int, orderNumber string) (*model.Order, error) {
	// order number Luhn validation
	if !luhn.Valid(orderNumber) {
		return nil, errors.Wrap(types.ErrUnprocessableEntity, fmt.Sprintf("luhn validation failed: %s", orderNumber))
	}

	orderCheck, _ := svc.store.Order.GetByNumber(ctx, orderNumber)
	if orderCheck != nil {
		if orderCheck.UserID == userID {
			return nil, errors.Wrap(types.StatusOK, fmt.Sprintf("duplicate: %s", orderNumber))
		}
		return nil, errors.Wrap(types.ErrConflict, fmt.Sprintf("duplicate: %s", orderNumber))
	}

	order := &model.Order{
		UserID:     userID,
		Status:     model.OrderStatusNew,
		Number:     orderNumber,
		UploadedAt: time.Now(),
	}

	result, err := svc.store.Order.Create(ctx, order)
	if err != nil {
		return nil, errors.Wrap(types.ErrDuplicateEntry, fmt.Sprintf("duplicate: %s", orderNumber))
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
