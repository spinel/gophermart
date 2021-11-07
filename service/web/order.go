package web

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spinel/gophermart/model"
	"github.com/spinel/gophermart/pkg/luhn"
	"github.com/spinel/gophermart/service/ext"
	"github.com/spinel/gophermart/store"
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
	orderService.workerUpdateStatus(5)

	return orderService

}

func (svc OrderWebService) workerUpdateStatus(interval int) {
	ctx := context.Background()

	go func() {
		for now := range time.Tick(time.Second * time.Duration(interval)) {
			fmt.Println(now)

			newOrders, err := svc.store.Order.GetByStatus(ctx, model.OrderStatusNew)
			if err != nil {
				log.Fatalf("error while load new orders: %s", err)
			}

			for _, order := range newOrders {
				orderResp, err := svc.ext.OrderStatus(ctx, order.Number)
				if err != nil {
					log.Fatalf("external api error: %s", err)
				}
				fmt.Println(">>> ", order.Number)
				fmt.Println(orderResp)
			}

		}
	}()

	fmt.Println("YES!")
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