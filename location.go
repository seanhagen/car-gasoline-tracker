package main

import (
	"database/sql"
	"fmt"
	"github.com/gchaincl/dotsql"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"strconv"
)

type Location struct {
	UUID      string
	Address   string // street address
	Longitude float64
	Latitude  float64
	Name      string // "brand" ?
	Visits    uint32 // how many times this location has been visited
	_saved    bool
	_changed  bool
}

func loadLocation(db *sql.DB, dot *dotsql.DotSql, address string) (*Location, error) {
	address += "%"

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

func searchLocation(db *sql.DB, dot *dotsql.DotSql) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		loc, err := loadLocation(db, dot, ps.ByName("address"))
		geo := getAddressGeo(ps.ByName("address"))

		if err != nil {
			fmt.Printf("Error loading location: %#v\n", err)
			panic(err)
		}

		io.WriteString(w, loc.Name)
		io.WriteString(w, "\n\n")
		io.WriteString(w, strconv.FormatFloat(geo.Lat, 'f', -1, 64))
		io.WriteString(w, ", ")
		io.WriteString(w, strconv.FormatFloat(geo.Lng, 'f', -1, 64))
	}
}

// geo := getAddressGeo("5036 Hastings St, Burnaby, BC")
// fmt.Printf("what: %#v\n\n", geo)
