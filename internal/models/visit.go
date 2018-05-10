package models

import (
	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Visit struct {
	ID         uuid.UUID
	LocationID uuid.UUID
	Odometer   int
	Liters     float32
	Cost       decimal.Decimal
}
