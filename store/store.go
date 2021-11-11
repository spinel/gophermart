package store

import (
	"fmt"
	"log"
	"time"

	"github.com/spinel/gophermart/config"
	"github.com/spinel/gophermart/logger"
	"github.com/spinel/gophermart/store/bun"
	"github.com/spinel/gophermart/store/memory"
)

// Store main struct
type Store struct {
	Bun      *bun.DB
	MemoryDB *memory.MemoryDB

	User        UserRepo
	Order       OrderRepo
	Preorder    PreorderRepo
	Transaction TransactionRepo
}

// New - create store
func New(cfg *config.Config) (*Store, error) {
	bunDB, err := bun.Dial(cfg)
	if err != nil {
		return nil, fmt.Errorf("pgdb.Dial failed: %w", err)
	}

	// Run migrations
	if bunDB != nil {
		log.Println("Running PostgreSQL migrations...")
		if err := runPgMigrations(); err != nil {
			return nil, fmt.Errorf("runPgMigrations failed: %w", err)
		}
	}

	var store Store
	store.MemoryDB = memory.New()

	// Init Postgres repositories
	if bunDB != nil {
		store.Bun = bunDB

		go store.KeepAlivePg(cfg)

		store.User = bun.NewUserRepo(bunDB)
		store.Order = bun.NewOrderRepo(bunDB)
		store.Preorder = bun.NewPreorderRepo(bunDB)
		store.Transaction = bun.NewTransactionPgRepo(bunDB)
	}

	return &store, nil
}

// KeepAlivePollPeriod is a Pg/MySQL keepalive check time period
const KeepAlivePollPeriod = 3

// KeepAlivePg makes sure PostgreSQL is alive and reconnects if needed
func (store *Store) KeepAlivePg(cfg *config.Config) {
	logger := logger.Get()
	var err error
	for {
		// Check if PostgreSQL is alive every 3 seconds
		time.Sleep(time.Second * KeepAlivePollPeriod)
		lostConnect := false
		if store.Bun == nil {
			lostConnect = true
		} else if _, err = store.Bun.Exec("SELECT 1"); err != nil {
			lostConnect = true
		}
		if !lostConnect {
			continue
		}
		logger.Debug().Msg("[store.KeepAlivePg] Lost PostgreSQL connection. Restoring...")
		store.Bun, err = bun.Dial(cfg)
		if err != nil {
			logger.Err(err)
			continue
		}
		logger.Debug().Msg("[store.KeepAlivePg] PostgreSQL reconnected")
	}
}
