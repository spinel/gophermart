package model

type Transaction struct {
	Base
	ID      int     `json:"id" pg:"id,notnull,pk"`
	UserID  int     `json:"user_id" pg:"number,notnull"`
	OrderID int     `json:"order_id" pg:"order_id,notnull"`
	Status  string  `json:"status"  pg:"status"`
	Amount  float64 `json:"amount,omitempty" pg:"amount"`
	Balance float64 `json:"balance" pg:"balance"`
}

type TransactionRequest struct {
	Order string  `json:"order"`
	Sum   float64 `json:"sum"`
}
