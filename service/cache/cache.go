package cache

import (
	"context"
	"fmt"

	"github.com/spinel/gophermart/store"
)

// UserWebService ...
type CacheService struct {
	ctx   context.Context
	store *store.Store
}

// NewCacheService is a cache service.
func NewCacheService(ctx context.Context, store *store.Store) *CacheService {
	return &CacheService{
		ctx:   ctx,
		store: store,
	}
}

func (cache CacheService) Do(ctx context.Context, commandName string, args ...interface{}) (reply interface{}, err error) {
	response, err := cache.store.Redis.Do(commandName, args...)
	if err != nil {
		// If there is an error in setting the cache, return an internal server error
		return nil, fmt.Errorf("svc.Cache session error: %w", err)
	}
	return response, nil
}
