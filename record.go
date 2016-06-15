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

// There are two ways to get a Record: create a brand new one, or
// load one from the database.

// NewRecord creates a new not-yet-saved-to-the-database Record
// that can be modified before saving.
func NewRecord(uuidGen func() string) *Record {
	if uuidGen == nil {
		uuidGen = uuid.NewV4().String
	}

	return &Record{
		uuidGenFunc: uuidGen,
	}
}

// LoadRecord attempts to load a Record from the database using the
// given UUID
func LoadRecord(uuid string) *Record {
	return &Record{}
}

func (rec *Record) build(r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(rec)
	if err != nil {
		return err
	}
	return nil
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
