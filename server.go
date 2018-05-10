package main

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	// "github.com/unrolled/render"
	"log"
	"net/http"
	"os"
	"strconv"
)

type (
	// AuthFunc is a function that can be used to authenticate a user. This is useful for apps that
	// need to validate against themselves ( ie, validating /deauth on the Auth service )
	AuthFunc func(http.Handler, http.ResponseWriter, *http.Request)

	// HandleFunc is a function that handles the incoming request
	HandleFunc func() httprouter.Handle

	// Route is a struct that provides the following information:
	// Uses AuthType to determine if the route needs to be authenticated before proceeding
	// If the AuthType is APPSELF, Auth is used to authenticate
	// If the request is authenticated, Handler is used to... handle the request
	Route struct {
		AuthType int
		Auth     AuthFunc
		Handler  HandleFunc
	}

	// RouteMap is the definition of HTTP verbs and routes that the application will handle
	// ex:
	//  routes := GetEmptyRoutes()
	//  routes["GET"]["/index] = CreateRoute(NONE, nil, someHandlerFunc)
	RouteMap map[string]map[string]Route
)

func properAuthType(t int) bool {
	for _, b := range authTypes {
		if b == t {
			return true
		}
	}
	return false
}

// CreateRoute returns a Route Struct
func CreateRoute(authtype int, auth AuthFunc, handler HandleFunc) Route {
	if !properAuthType(authtype) {
		authtype = NONE
	}

	return Route{
		AuthType: authtype,
		Auth:     auth,
		Handler:  handler,
	}
}

// GetEmptyRoutes returns an initialized and empty map to be populated with the routes
// to be passed into CreateServer
func GetEmptyRoutes() RouteMap {
	routes := make(map[string]map[string]Route)
	verbs := []string{"GET", "PUT", "POST", "OPTIONS", "PATCH", "DELETE"}
	for _, v := range verbs {
		routes[v] = make(map[string]Route)
	}
	return routes
}

// CreateServer returns a http.HandleFunc with all of the middleware for KloudKtrl
// already set up and initialized. It sets up the following middleware:
//  * CORS handler
//  * Sets up gorilla/context
//  * Sets up the custom token authentication middleware
//  * Sets up some logging middleware
func CreateServer(routes RouteMap) http.Handler {
	router := httprouter.New()
	setcontext := makeContextMiddleware()

	corHandler := cors.New(
		cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowCredentials: true,
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
			AllowedHeaders:   []string{"X-Auth-Token", "Accept", "Content-Type"},
			ExposedHeaders:   []string{"X-Auth-Token"},
		},
	)

	handlers := alice.New(
		setcontext,
		tokenAuth(routes),
		context.ClearHandler,
		logMiddleware,
		corHandler.Handler,
	)

	for verb, verbs := range routes {
		for route, part := range verbs {
			switch verb {
			case "GET":
				router.GET(route, part.Handler())
			case "POST":
				router.POST(route, part.Handler())
			case "PUT":
				router.PUT(route, part.Handler())
			case "DELETE":
				router.DELETE(route, part.Handler())
			case "HEAD":
				router.HEAD(route, part.Handler())
			case "OPTIONS":
				router.OPTIONS(route, part.Handler())
			}
		}
	}
	return handlers.Then(router)
}

func server() {
	// get database connection and queries
	//db, dot := getDbQuery()

	routes := GetEmptyRoutes()

	routes["GET"]["/"] = CreateRoute(NONE, nil, indexRoute)
	routes["POST"]["/oauth2"] = CreateRoute(NONE, nil, oauth2Route)
	routes["GET"]["/oauth2callback"] = CreateRoute(NONE, nil, oauth2CallbackRoute)

	routes["GET"]["/locations"] = CreateRoute(NONE, nil, locationsIndex)
	routes["GET"]["/locations/:id"] = CreateRoute(NONE, nil, locationsGet)
	routes["POST"]["/locations"] = CreateRoute(TOKEN, nil, locationsCreate)
	routes["PUT"]["/locations/:id"] = CreateRoute(TOKEN, nil, locationsUpdate)
	routes["DELETE"]["/locations/:id"] = CreateRoute(TOKEN, nil, locationsDelete)

	routes["GET"]["/records"] = CreateRoute(TOKEN, nil, recordsList)
	routes["GET"]["/records/:id"] = CreateRoute(TOKEN, nil, recordsGet)
	routes["POST"]["/records"] = CreateRoute(TOKEN, nil, recordsCreate)
	routes["PUT"]["/records/:id"] = CreateRoute(TOKEN, nil, recordsUpdate)
	routes["DELETE"]["/records/:id"] = CreateRoute(TOKEN, nil, recordsDelete)

	// start the server!
	router := CreateServer(routes)

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = strconv.Itoa(*serverPortFlag)
	}

	listen := ":" + port
	fmt.Printf("Starting server on %#v\n", listen)
	log.Fatal(http.ListenAndServe(listen, router))
}
