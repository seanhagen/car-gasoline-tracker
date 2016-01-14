package main

import (
	"database/sql"
	"github.com/gchaincl/dotsql"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func server(db *sql.DB, dot *dotsql.DotSql) {
	router := httprouter.New()

	//router.GET("/locations", listLocations(db, dot))
	//router.GET("/locations/:id", fetchLocation(db, dot))
	router.GET("/locations/search/:address", searchLocation(db, dot))

	//router.POST("/locations", createLocation(db, dot))

	log.Fatal(http.ListenAndServe(":8000", router))
}
