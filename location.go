package main

import (
	"bytes"
	"database/sql"
)

type Location struct {
	GUID string
	Address string // street address
	Longitude float32
	Latitude float32
	Name string // "brand" ?
	Visits uint32 // how many times this location has been visited
}

func loadLocation(db *sql.DB, address string) (*Location,error) {
	var query bytes.Buffer
	query.WriteString("select id, name, longitude, latitude, address, visits from locations ")
	query.WriteString("where address like '")
	query.WriteString(address)
	query.WriteString("%'")

	row := db.QueryRow(query.String())

	var loc Location

	err := row.Scan(&loc.GUID, &loc.Name, &loc.Longitude, &loc.Latitude, &loc.Address, &loc.Visits)

	switch {
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, err
	}

	return &loc, nil
}
