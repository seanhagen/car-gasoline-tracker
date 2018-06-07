package models

import (
	"time"

	postgis "github.com/cridenour/go-postgis"
	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

// Station a fillup station
type Station struct {
	ID        uuid.UUID       `json:"id" db:"id"`
	Name      string          `json:"name" db:"name"`
	Address   string          `json:"address" db:"address"`
	Location  postgis.PointZ  `json:"location" db:"location"` // Y is Latitude, X is Longitude
	Distance  decimal.Decimal `json:"distance,omitempty" db:"distance"`
	Visits    int             `json:"visits" db:"visits"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt time.Time       `json:"updated_at" db:"updated_at"`
}
