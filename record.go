package main

import (
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/satori/go.uuid"
	"net/http"
)

type Record struct {
	UUID         string
	LocationUUID string  // location of the gas station
	Odometer     uint32  // odometer reading when gas purchased
	Liters       float32 // amount of gas purchased
	Cost         uint16  // cost in cents
}

func (rec *Record) build(r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(rec)
	if err != nil {
		return err
	}
	return nil
}

func (rec *Record) create(r *http.Request) error {
	ce := context.Get(r, "extras")

	rec.UUID = uuid.NewV4().String()

	_, err := ce.(Extra).dot.Exec(
		ce.(Extra).db,
		"create-notification",
		rec.UUID,
		rec.LocationUUID,
		rec.Odometer,
		rec.Liters,
		rec.Cost,
	)

	return err
}

func (rec *Record) update(r *http.Request) error {
	return nil
}

func (rec *Record) fetch(uuid string, r *http.Request) error {
	ce := context.Get(r, "extras")
	row, err := ce.(Extra).dot.QueryRow(
		ce.(Extra).db,
		"fetch-notification",
		uuid,
	)

	row.Scan(&rec.LocationUUID)

	return err
}
