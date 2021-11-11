package bun

import (
	"context"

	"github.com/spinel/gophermart/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

// OrderPgRepo ...
type OrderRepo struct {
	db *DB
}

// NewOrderPgRepo ...
func NewOrderRepo(db *DB) *OrderRepo {
	orderRepo := &OrderRepo{
		db: db,
	}

	return orderRepo
}

// Create creates new order in Postgres.
func (repo *OrderRepo) Create(ctx context.Context, order *model.Order) (*model.Order, error) {
	_, err := repo.db.NewInsert().
		Model(order).
		Returning("*").
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return order, nil
}

// Update order in Postgres.
func (repo *OrderRepo) Update(ctx context.Context, order *model.Order) error {
	_, err := repo.db.NewUpdate().
		Model(order).
		WherePK().
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

// List of orders of current user.
func (repo *OrderRepo) List(ctx context.Context, userID int) ([]model.Order, error) {
	var orders []model.Order
	err := repo.db.NewSelect().
		Model(&orders).
		Where("? = ?", bun.Ident("user_id"), userID).
		Where(notDeleted).
		Order("uploaded_at ASC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return orders, nil
}

// GetByStatus orders list.
func (repo *OrderRepo) GetByStatus(ctx context.Context, orderStatus string) ([]model.Order, error) {
	var orders []model.Order
	err := repo.db.NewSelect().
		Model(&orders).
		Where("? = ?", bun.Ident("status"), orderStatus).
		Where(notDeleted).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return orders, nil
}

// GetByNumber orders list.
func (repo *OrderRepo) GetByNumber(ctx context.Context, orderNumber string) (*model.Order, error) {
	var order model.Order
	err := repo.db.NewSelect().
		Model(&order).
		Where("? = ?", bun.Ident("number"), orderNumber).
		//Where("? = ?", bun.Ident("user_id"), userID).
		Where(notDeleted).
		Scan(ctx)

	if err != nil {
		pqErr := err.(pgdriver.Error)
		if pqErr.IntegrityViolation() { // duplicate
			return nil, nil
		}
		return nil, err
	}

	return &order, nil
}
