package main

import (
	"database/sql"
	"github.com/gchaincl/dotsql"
	_ "github.com/lib/pq"
)

type Database struct {
	dbq  dotsql.Queryer
	dbqr dotsql.QueryRower
	dot  *dotsql.DotSql
}

func (d *Database) Query(name string, args ...interface{}) (*sql.Rows, error) {
	return d.dot.Query(d.dbq, name, args)
}

func (d *Database) QueryRow(name string, args ...interface{}) (*sql.Row, error) {
	return d.dot.QueryRow(d.dbqr, name, args)
}

func getDbQuery() *Database {
	connString := getServiceURI("elephantsql")
	if len(connString) == 0 {
		connString = *databaseConnectionString
	}

	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	// load up queries ( dotsql )
	dot, err := dotsql.LoadFromFile("./db/queries.sql")
	if err != nil {
		panic(err)
	}

	return &Database{
		dbq:  db,
		dbqr: db,
		dot:  dot,
	}
}
