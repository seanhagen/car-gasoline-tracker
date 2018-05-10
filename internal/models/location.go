package models

import (
	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Location struct {
	ID        uuid.UUID
	Address   string
	Latitude  decimal.Decimal
	Longitude decimal.Decimal
	Name      string
	Visits    int
}
