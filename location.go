package main

import (
	"database/sql"
	"fmt"
	"github.com/gchaincl/dotsql"
)

type Location struct {
	GUID      string
	Address   string // street address
	Longitude float32
	Latitude  float32
	Name      string // "brand" ?
	Visits    uint32 // how many times this location has been visited
}

func loadLocation(db *sql.DB, dot *dotsql.DotSql, address string) (*Location, error) {
	address += "%"
	fmt.Printf("address: %#v\n", address)

	row, err := dot.QueryRow(db, "fetch-location", address)
	if err != nil {
		panic(err)
	}

	var loc Location

	err = row.Scan(&loc.GUID, &loc.Name, &loc.Longitude, &loc.Latitude, &loc.Address, &loc.Visits)

	switch {
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, err
	}

	return &loc, nil
}
