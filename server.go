package main

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	"github.com/unrolled/render"
	"log"
	"net/http"
	"os"
	"strconv"
)

func server() {
	// get cloudfoundry env

	// appEnv, _ := cfenv.Current()

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

	router.GET("/", index())

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

	var addr string
	var port string

	if addr = os.Getenv("HOST"); len(addr) == 0 {
		addr = *serverAddressFlag
	}
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = strconv.Itoa(*serverPortFlag)
	}

	listen := addr + ":" + port

	fmt.Printf("Starting server on %#v\n", listen)

	log.Fatal(http.ListenAndServe(listen, handlers.Then(router)))
}
