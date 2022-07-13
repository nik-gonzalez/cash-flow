package dto

import "cash-flow-service/customType"

type CashFlow struct {
	ID      string                  `json:"id"`
	Name    string                  `json:"name"`
	Remarks string                  `json:"remarks"`
	Amount  float64                 `json:"amount"`
	Date    customType.JsonDateTime `json:"date"`
}
