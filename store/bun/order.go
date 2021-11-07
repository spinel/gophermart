package bun

import (
	"context"

	"github.com/spinel/gophermart/model"
	"github.com/uptrace/bun"
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
