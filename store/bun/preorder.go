package bun

import (
	"context"

	"github.com/spinel/gophermart/model"
)

// PreorderRepo ...
type PreorderRepo struct {
	db *DB
}

// NewOrderPgRepo ...
func NewPreorderRepo(db *DB) *PreorderRepo {
	preorderRepo := &PreorderRepo{
		db: db,
	}

	return preorderRepo
}

// Create creates new preorder in Postgres.
func (repo *PreorderRepo) Create(ctx context.Context, preorder *model.Preorder) (*model.Preorder, error) {
	_, err := repo.db.NewInsert().
		Model(preorder).
		Returning("*").
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return preorder, nil
}
