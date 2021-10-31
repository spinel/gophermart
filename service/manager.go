package service

import (
	"context"
	"errors"

	"github.com/spinel/gophermart/service/cache"
	"github.com/spinel/gophermart/service/web"
	"github.com/spinel/gophermart/store"
)

// Manager is just a collection of all services we have in the project
type Manager struct {
	User  UserService
	Order OrderService
	Cache CacheService
}

// NewManager creates new service manager
func NewManager(ctx context.Context, store *store.Store) (*Manager, error) {
	if store == nil {
		return nil, errors.New("No store provided")
	}
	return &Manager{
		User:  web.NewUserWebService(ctx, store),
		Order: web.NewOrderWebService(ctx, store),
		Cache: cache.NewCacheService(ctx, store),
	}, nil
}
