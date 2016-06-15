package main

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

var index = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

func indexRoute() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		index.Execute(w, nil)
	}
}

func oauth2Route() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ce := context.Get(r, "extras")
		retval := SkeletonMessage{Message: "This should actually redirect to google for oauth..."}
		ce.(Extra).render.JSON(w, http.StatusOK, retval)
	}
}

func oauth2CallbackRoute() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ce := context.Get(r, "extras")
		retval := SkeletonMessage{Message: "This should actually log a user in..."}
		ce.(Extra).render.JSON(w, http.StatusOK, retval)
	}
}
