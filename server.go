package main

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	//"github.com/unrolled/render"
	"log"
	"net/http"
	"os"
	"strconv"
)

func server() {
	mq := getMQConn()
	defer mq.Close()

	ch, err := mq.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	router := httprouter.New()
	setcontext := makeSetContextMiddleware(ch, queue)

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

	router.ServeFiles("/static/*filepath", http.Dir("static"))
	router.GET("/", indexRoute())
	router.GET("/appenv", appenvRoute())
	router.GET("/queue/:word", postToQueue())
	router.GET("/ws", websocketRoute())

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

	listen := ":" + port // addr + ":" + port

	fmt.Printf("Starting server on %#v\n", listen)

	log.Fatal(http.ListenAndServe(listen, handlers.Then(router)))
}
