package main

import (
	"encoding/json"
	//"github.com/gorilla/context"
	"github.com/satori/go.uuid"
	"net/http"
)

type Record struct {
	UUID         string  `json:"uuid"`
	LocationUUID string  `json:"location_uuid"` // location of the gas station
	Odometer     uint32  `json:"odometer"`      // odometer reading when gas purchased
	Liters       float32 `json:"liters"`        // amount of gas purchased
	Cost         uint16  `json:"cost"`          // cost in cents

	//private!
	uuidGenFunc func() string
}

// The ways to get a Record: create from scratch, build from request, load from db.

// NewRecord creates a new not-yet-saved-to-the-database Record
// that can be modified before saving.
func NewRecord(uuidGen func() string) (*Record, error) {
	if uuidGen == nil {
		uuidGen = uuid.NewV4().String
	}

	return &Record{
		uuidGenFunc: uuidGen,
	}, nil
}

func BuildRecord(r *http.Request) (rec *Record, err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(rec)
	return
}

// LoadRecord attempts to load a Record from the database using the
// given UUID
func LoadRecord(uuid string) (rec *Record, err error) {
	return &Record{}, nil
}

func (rec *Record) Create(r *http.Request) error {
	// ce := context.Get(r, "extras")

	rec.UUID = rec.uuidGenFunc()

	// _, err := ce.(Extra).dot.Exec(
	// 	ce.(Extra).db,
	// 	"create-notification",
	// 	rec.UUID,
	// 	rec.LocationUUID,
	// 	rec.Odometer,
	// 	rec.Liters,
	// 	rec.Cost,
	// )

	return nil
}

func (rec *Record) update(r *http.Request) error {
	return nil
}

func (rec *Record) fetch(uuid string, r *http.Request) error {
	// ce := context.Get(r, "extras")
	// row, err := ce.(Extra).dot.QueryRow(
	// 	ce.(Extra).db,
	// 	"fetch-notification",
	// 	uuid,
	// )

	// row.Scan(&rec.LocationUUID)

	return nil
}
