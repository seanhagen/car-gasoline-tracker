package main

import (
	// "database/sql"
	// "github.com/gorilla/context"
	"net/http"
)

type Record struct {
	UUID       string
	LocationId string  // location of the gas station
	Odometer   uint32  // odometer reading when gas purchased
	Liters     float32 // amount of gas purchased
	Cost       uint16  // cost in cents
}

func (rec *Record) create(r *http.Request) error {
	return nil
}

func (rec *Record) save(r *http.Request) error {
	return nil
}

func (rec *Record) fetch(uuid string, r *http.Request) error {
	return nil
}
