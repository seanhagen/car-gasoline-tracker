package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gchaincl/dotsql"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

type EnvConfig struct {
	Confpath string
}

var conf EnvConfig

func main() {
	err := envconfig.Process("gasapp", &conf)
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	fmt.Printf("envconfig: %#v\n\n", conf)

	db, err := sql.Open("postgres", loadDBConnectionString())
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	dot, err := dotsql.LoadFromFile("db/queries.sql")
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	server(db, dot)
}
