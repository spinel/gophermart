package service

import (
	"context"
	"fmt"

	"github.com/spinel/gophermart/service/ext"
	"github.com/spinel/gophermart/service/memory"
	"github.com/spinel/gophermart/service/web"
	"github.com/spinel/gophermart/store"
)

// Manager is just a collection of all services we have in the project
type Manager struct {
	User        UserService
	Order       OrderService
	Transaction TransactionService
	Memory      MemoryService
}

// NewManager creates new service manager
func NewManager(ctx context.Context, store *store.Store) (*Manager, error) {
	if store == nil {
		return nil, fmt.Errorf("no store provided")
	}

	extService := ext.NewOrderWebService(ctx, "http://localhost:8080")

	return &Manager{
		User:        web.NewUserWebService(ctx, store),
		Order:       web.NewOrderWebService(ctx, store, extService),
		Transaction: web.NewTransactionWebService(ctx, store),
		Memory:      memory.NewMemoryService(ctx, store),
	}, nil
}
