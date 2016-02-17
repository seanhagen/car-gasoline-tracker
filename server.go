package main

import (
	"fmt"
	"github.com/seanhagen/gas-web/Godeps/_workspace/src/github.com/gorilla/context"
	"github.com/seanhagen/gas-web/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/seanhagen/gas-web/Godeps/_workspace/src/github.com/justinas/alice"
	"github.com/seanhagen/gas-web/Godeps/_workspace/src/github.com/rs/cors"
	"github.com/seanhagen/gas-web/Godeps/_workspace/src/github.com/unrolled/render"
	"log"
	"net/http"
	"strconv"
)

func server() {
	// get database connection and queries
	db, dot := getDbQuery()

	router := httprouter.New()
	render := render.New(render.Options{
		IndentJSON: true,
	})
	setcontext := makeSetContextMiddleware(render, db, dot)

	corHandler := cors.New(
		cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowCredentials: true,
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
			AllowedHeaders:   []string{"X-Auth-Token", "Accept", "Content-Type"},
			ExposedHeaders:   []string{"X-Auth-Token"},
		},
	)

	handlers := alice.New(setcontext, TokenAuth, context.ClearHandler, Log, corHandler.Handler)

	router.GET("/locations", locationsList())
	router.GET("/locations/:id", locationsFetch())
	router.POST("/locations", locationsCreate())
	router.PUT("/locations/:id", locationsUpdate())
	router.DELETE("/locations/:id", locationsDelete())

	router.GET("/records", recordsList())
	router.GET("/records/:id", recordsFetch())
	router.POST("/records", recordsCreate())
	router.PUT("/records/:id", recordsUpdate())
	router.DELETE("/records/:id", recordsDelete())

	port := ":" + strconv.Itoa(*serverPortFlag)
	fmt.Printf("Starting server on port %#v\n", port)
	log.Fatal(http.ListenAndServe(port, handlers.Then(router)))
}
