package store

import (
	"github.com/shopspring/decimal"
	"time"
)

type Entity struct {
	ID         string           `json:"id"`
	IsActive   *bool            `json:"is_active"`
	MerchantId *string          `json:"merchant_id"`
	Name       *string          `json:"name"`
	Location   *string          `json:"location"`
	Rating     *decimal.Decimal `json:"rating"`
	CurrencyId *string          `json:"currency_id"`
	CityId     *string          `json:"city_id"`
	ScheduleId *string          `json:"schedule_id"`
	CreatedAt  time.Time        `json:"-" db:"created_at"`
	UpdatedAt  time.Time        `json:"-" db:"updated_at"`
}

type Cities struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Geocenter string `json:"geocenter"`
}

type Schedule struct {
	ID       string `json:"id"`
	IsActive bool   `json:"is_active"`
	Periods  Period
}

type Delivery struct {
	ID       string `json:"id"`
	IsActive bool   `json:"is_active"`
	Periods  Period `json:"periods"`
	Areas    Area   `json:"area"`
}

type Period struct {
	Day  string `json:"day"`
	From string `json:"from"`
	To   string `json:"to"`
}

type Area struct {
	From string `json:"from"`
	To   string `json:"to"`
}
