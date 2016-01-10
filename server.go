package main

import (
	"io"
	"fmt"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"database/sql"

	_ "github.com/lib/pq"
)

func test(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
		loc, err := loadLocation(db, "5036 Hasting")

		if err != nil {
			fmt.Printf("Error loading location: %#v\n", err)
			panic(err)
		}

		io.WriteString(w, loc.Name)
		// fmt.Printf("%#v\n", loc)
	}
}

func testParam(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		loc, err := loadLocation(db, ps.ByName("address"))

		if err != nil {
			fmt.Printf("Error loading location: %#v\n", err)
			panic(err)
		}

		io.WriteString(w, loc.Name)
	}
}

func server(db *sql.DB) {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", test(db))
	// http.ListenAndServe(":8000", mux)

	router := httprouter.New()
	router.GET("/", test(db))
	router.GET("/test/:address", testParam(db))

	log.Fatal(http.ListenAndServe(":8000", router))
}
