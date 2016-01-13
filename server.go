package main

import (
	"database/sql"
	"fmt"
	"github.com/gchaincl/dotsql"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
)

func test(db *sql.DB, dot *dotsql.DotSql) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		loc, err := loadLocation(db, dot, "5036 Hasting")

		if err != nil {
			fmt.Printf("Error loading location: %#v\n", err)
			panic(err)
		}

		io.WriteString(w, loc.Name)
		fmt.Printf("%#v\n", loc)
	}
}

func testParam(db *sql.DB, dot *dotsql.DotSql) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		loc, err := loadLocation(db, dot, ps.ByName("address"))

		if err != nil {
			fmt.Printf("Error loading location: %#v\n", err)
			panic(err)
		}

		io.WriteString(w, loc.Name)
	}
}

func server(db *sql.DB, dot *dotsql.DotSql) {
	router := httprouter.New()
	router.GET("/", test(db, dot))
	router.GET("/test/:address", testParam(db, dot))

	log.Fatal(http.ListenAndServe(":8000", router))
}
