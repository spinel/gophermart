package model

const (
	TransactionTypeRefill   = "refill"
	TransactionTypeWithdraw = "withdraw"
)

type Transaction struct {
	Base
	ID      int     `json:"id" pg:"id,notnull,pk"`
	UserID  int     `json:"user_id" pg:"number,notnull"`
	OrderID int     `json:"order_id" pg:"order_id,notnull"`
	Type    string  `json:"type"  pg:"type"`
	Amount  float64 `json:"amount,omitempty" pg:"amount"`
}

type TransactionRequest struct {
	Order string  `json:"order"`
	Sum   float64 `json:"sum"`
}
