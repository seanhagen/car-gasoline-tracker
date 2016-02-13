package main

var goticFiles map[string]string

func init() {
  goticFiles = make(map[string]string)

  goticFiles["./db/queries.sql"] = "-- name: fetch-location\nSELECT id,name,longitude,latitude,address,visits FROM locations WHERE address LIKE $1\n"

}
