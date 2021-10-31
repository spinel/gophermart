package model

import (
	"context"
	"time"
)

// Models hold registered models in-memory
var Models []interface{}

// Base contains common fields for all tables
type Base struct {
	CreatedAt time.Time  `json:"created_at,omitempty" pg:"created_at,notnull"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" pg:"updated_at,notnull"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" pg:"deleted_at"`
}

// BeforeInsert hooks into insert operations, setting createdAt and updatedAt to current time
func (b *Base) BeforeInsert(ctx context.Context) (context.Context, error) {
	now := time.Now().UTC()
	if b.CreatedAt.IsZero() {
		b.CreatedAt = now
	}
	if b.UpdatedAt.IsZero() {
		b.UpdatedAt = now
	}
	return ctx, nil
}

// BeforeUpdate hooks into update operations, setting updatedAt to current time
func (b *Base) BeforeUpdate(ctx context.Context) (context.Context, error) {
	b.UpdatedAt = time.Now().UTC()
	return ctx, nil
}

// Delete sets deleted_at time to current_time
func (b *Base) Delete() {
	t := time.Now().UTC()
	b.DeletedAt = &t
}

// Register is used for registering models
func Register(m interface{}) {
	Models = append(Models, m)
}
