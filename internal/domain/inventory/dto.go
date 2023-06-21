package inventory

import "github.com/shopspring/decimal"

type Request struct {
	StoreId      string          `json:"store_id" validate:"required"`
	ProductId    string          `json:"product_id" validate:"required"`
	Quantity     uint            `json:"quantity" validate:"required"`
	Price        decimal.Decimal `json:"price" validate:"required"`
	PriceSpecial decimal.Decimal `json:"price_special"`
	IsAvailable  bool            `json:"is_available"`
}

type Response struct {
	ID            string          `json:"id"`
	StoreId       string          `json:"store_id"`
	ProductId     string          `json:"product_id"`
	Quantity      uint            `json:"quantity"`
	QuantityMin   uint            `json:"quantity_min"`
	QuantityMax   uint            `json:"quantity_max"`
	Price         decimal.Decimal `json:"price"`
	PriceSpecial  decimal.Decimal `json:"price_special"`
	PricePrevious decimal.Decimal `json:"price_previous"`
	IsAvailable   bool            `json:"is_available"`
}
