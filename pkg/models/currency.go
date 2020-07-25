package models

import "github.com/mewben/arrstate/internal/enums"

// CurrencyModel -
type CurrencyModel struct {
	Currency string `bson:"currency" json:"currency"`
}

// NewCurrencyModel -
func NewCurrencyModel() CurrencyModel {
	return CurrencyModel{
		Currency: enums.DefaultCurrency,
	}
}
