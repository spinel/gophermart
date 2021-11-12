package bun

import (
	"context"

	"github.com/spinel/gophermart/model"
	"github.com/uptrace/bun/driver/pgdriver"
)

// UserPgRepo ...
type UserPgRepo struct {
	db *DB
}

// NewUserRepo ...
func NewUserRepo(db *DB) *UserPgRepo {
	return &UserPgRepo{db: db}
}

// Create creates User in Postgres
func (repo *UserPgRepo) Create(ctx context.Context, user *model.User) (*model.User, error) {
	_, err := repo.db.NewInsert().
		Model(user).
		Returning("*").
		Exec(ctx)

	if err != nil {
		pqErr := err.(pgdriver.Error)
		if pqErr.IntegrityViolation() { // duplicate
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

// GetByLogin retrieve user by login.
func (repo *UserPgRepo) GetByLogin(ctx context.Context, login string) (*model.User, error) {
	var user model.User
	err := repo.db.NewSelect().
		Model(&user).
		Where("login = ?", login).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
