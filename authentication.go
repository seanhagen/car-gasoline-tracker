// Middleware to handle validating a the token in the X-Auth-Token header.
// Currently doesn't do anything because the Auth app isn't up yet. Need to get
// that deployed so this can be fleshed out.
package main

import (
	"fmt"
	"github.com/gorilla/context"
	"net/http"
)

type user struct {
	UUID     string
	Username string
}

var (
	// NONE No authentication needed for a route
	NONE = 1 << 0
	// TOKEN Checks the X-Auth-Token for a valid OAuth token
	TOKEN = 1 << 1
	// SUPERUSER Checks the user against the superuser method
	SUPERUSER = 1 << 2
)

var authTypes = map[string]int{
	"none":  NONE,
	"token": TOKEN,
	"su":    SUPERUSER,
}

func tokenAuth(routes RouteMap) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			route := routes[r.URL.Path][r.Method]
			fmt.Printf("Got route: %#v ( url: %#v )", route, r.URL)

			switch route.AuthType {
			case NONE:
				h.ServeHTTP(w, r)
				return
			case TOKEN:
				authToken(h, w, r)
				return
			case SUPERUSER:
				superuserAuth(h, w, r)
				return
			}

			// continue down middleware chain
			h.ServeHTTP(w, r)
		})
	}
}

// superuserAuth validates the user against the Auth service /su route
func superuserAuth(h http.Handler, w http.ResponseWriter, r *http.Request) {
	h.ServeHTTP(w, r)
}

// authUserToken ...
func authToken(h http.Handler, w http.ResponseWriter, r *http.Request) {
	// get token from header
	auth := r.Header.Get("X-Auth-Token")

	// if token doesn't exist, set header to 400 ( or whatever the code for 'unauthorized' is )
	if auth == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// fetch the user information from auth service based on token

	// find or create user
	u := user{
		UUID:     "2a7c7819-39c3-4c19-9479-3e69db522802",
		Username: "shagen",
	}

	// set 'user' context to that user
	context.Set(r, "user", u)

	h.ServeHTTP(w, r)
}
