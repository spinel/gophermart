package memory

import (
	"context"

	"github.com/spinel/gophermart/store"
)

// MemoryService ...
type MemoryService struct {
	ctx   context.Context
	store *store.Store
}

// NewCacheService is a cache service.
func NewMemoryService(ctx context.Context, store *store.Store) *MemoryService {
	return &MemoryService{
		ctx:   ctx,
		store: store,
	}
}

func (m MemoryService) Add(ctx context.Context, sessionToken string, userID int) {
	m.store.MemoryDB.Add(sessionToken, userID)
}

func (m MemoryService) Get(ctx context.Context, sessionToken string) int {
	return m.store.MemoryDB.Get(sessionToken)
}
