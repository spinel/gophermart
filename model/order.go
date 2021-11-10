package model

import (
	"time"
)

const (
	OrderStatusNew        = "NEW"
	OrderStatusProcessing = "PROCESSING"
	OrderStatusInvalid    = "INVALID"
	OrderStatusProcessed  = "PROCESSED"
)

// Order is a shop order.
type Order struct {
	Base
	ID         int       `json:"id" pg:"id,notnull,pk"`
	Number     int       `json:"number" validate:"required" pg:"number,unique,notnull"`
	UserID     int       `json:"user_id" pg:"number,notnull"`
	Status     string    `json:"status"  pg:"status"`
	Accural    float64   `json:"accural,omitempty" pg:"accural"`
	UploadedAt time.Time `json:"uploaded_at" pg:"uploaded_at,notnull"`
}

type OrderAccural struct {
	OrderID int
	Number  string
}

//ExtOrder is an order response of accural system.
type ExtOrder struct {
	Order   string  `json:"order"`
	Status  string  `json:"status"`
	Accural float64 `json:"accural"`
}
