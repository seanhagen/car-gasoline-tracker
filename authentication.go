// Middleware to handle validating a the token in the X-Auth-Token header.
// Currently doesn't do anything because the Auth app isn't up yet. Need to get
// that deployed so this can be fleshed out.
package main

import (
	"github.com/gorilla/context"
	"net/http"
)

func TokenAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get token from header
		//auth := r.Header.Get("X-Auth-Token")

		// if token doesn't exist, set header to 400 ( or whatever the code for 'unauthorized' is )

		// fetch the user information from auth service based on token

		// find or create user
		u := User{
			UUID:     "2a7c7819-39c3-4c19-9479-3e69db522802",
			Username: "shagen",
		}

		// set 'user' context to that user
		context.Set(r, "user", u)

		// continue down middleware chain
		h.ServeHTTP(w, r)
	})
}
