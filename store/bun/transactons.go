package bun

import (
	"context"

	"github.com/spinel/gophermart/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

// TransactionPgRepo ...
type TransactionPgRepo struct {
	db *DB
}

// NewTransactionPgRepo ...
func NewTransactionPgRepo(db *DB) *TransactionPgRepo {
	return &TransactionPgRepo{db: db}
}

// Create transaction in Postgres
func (repo *TransactionPgRepo) Create(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	_, err := repo.db.NewInsert().
		Model(transaction).
		Returning("*").
		Exec(ctx)

	if err != nil {
		pqErr := err.(pgdriver.Error)
		if pqErr.IntegrityViolation() { // duplicate
			return nil, nil
		}
		return nil, err
	}

	return transaction, nil
}

// Withdraw transaction.
func (repo *TransactionPgRepo) Withdraw(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	return transaction, nil
}

// Balance of current user.
func (repo *TransactionPgRepo) Balance(ctx context.Context, userID int) (float64, error) {
	var b model.Transaction
	err := repo.db.NewSelect().
		Model(&b).
		ColumnExpr("SUM(amount) AS amount").
		Where("? = ?", bun.Ident("user_id"), userID).
		Where(notDeleted).
		Scan(ctx)

	if err != nil {
		return 0, err
	}

	return b.Amount, nil
}
