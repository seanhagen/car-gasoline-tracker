package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db sql.DB

func main() {
	db, err := sql.Open("postgres", loadDBConnectionString())

	if err != nil {
		panic(err)
	}

	server(db)
}
