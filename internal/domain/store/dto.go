package store

import (
	"github.com/shopspring/decimal"
	"time"
)

type Request struct {
}

type Response struct {
	ID              string          `json:"id"`
	IsActive        bool            `json:"is_active"`
	Name            string          `json:"name"`
	Location        string          `json:"location"`
	Rating          decimal.Decimal `json:"rating"`
	CurrencyId      string          `json:"currency_id"`
	City            Cities          `json:"city"`
	ScheduleRequest Schedule        `json:"schedule_id"`
	CreatedAt       time.Time       `json:"-" db:"created_at"`
	UpdatedAt       time.Time       `json:"-" db:"updated_at"`
}
