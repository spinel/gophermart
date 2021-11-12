package model

type Preorder struct {
	Base
	ID     int     `json:"-" pg:"id,notnull,pk"`
	Number  string  `json:"number" validate:"required" pg:"number,unique,notnull"`
	Amount float64 `json:"amount" pg:"amount"`
}
