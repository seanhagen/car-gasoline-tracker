package models

import (
	"time"

	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

// Fillup is a record of filling up the tank
type Fillup struct {
	ID       uuid.UUID       `json:"id"`
	Cost     decimal.Decimal `json:"cost"`
	Currency string          `json:"currency"`
	Amount   decimal.Decimal `json:"amount"`
	Type     string          `json:"type"`
	Date     time.Time       `json:"date"`
}
