package main

import (
	"database/sql"
	"github.com/seanhagen/gas-web/Godeps/_workspace/src/github.com/gchaincl/dotsql"
	_ "github.com/seanhagen/gas-web/Godeps/_workspace/src/github.com/lib/pq"
)

func getDbQuery() (*sql.DB, *dotsql.DotSql) {
	db, err := sql.Open("postgres", *databaseConnectionString)
	if err != nil {
		panic(err)
	}

	// load up queries ( dotsql )
	dot, err := dotsql.LoadFromString(goticFiles["./db/queries.sql"])
	if err != nil {
		panic(err)
	}

	return db, dot
}
