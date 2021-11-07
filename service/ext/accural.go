package ext

import (
	"context"
	"fmt"

	"github.com/spinel/gophermart/model"
)

// ExtAccuralService ...
type ExtAccuralService struct {
	ctx     context.Context
	address string
}

// ExtAccuralService is an external accural service builder.
func NewOrderWebService(ctx context.Context, address string) *ExtAccuralService {
	orderService := &ExtAccuralService{
		ctx:     ctx,
		address: address,
	}

	return orderService
}

// OrderStatus external status responseof accural system.
func (ext ExtAccuralService) OrderStatus(ctx context.Context, number int) (*model.ExtOrder, error) {
	var order *model.ExtOrder
	fmt.Println(fmt.Sprintf(ext.address, "/api/orders"))

	return order, nil
}
