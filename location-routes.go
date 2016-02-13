package main

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func locationsCreate() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ce := context.Get(r, "extras")
		msg := SkeletonMessage{Message: "Yup, created!"}
		ce.(Extra).render.JSON(w, http.StatusOK, msg)
	}
}

func locationsList() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ce := context.Get(r, "extras")
		var list []SkeletonMessage = []SkeletonMessage{
			SkeletonMessage{Message: "This is a thing"},
			SkeletonMessage{Message: "This is a thing"},
			SkeletonMessage{Message: "This is a thing"},
			SkeletonMessage{Message: "This is a thing"},
			SkeletonMessage{Message: "This is a thing"},
			SkeletonMessage{Message: "This is a thing"},
			SkeletonMessage{Message: "This is a thing"},
		}
		ce.(Extra).render.JSON(w, http.StatusOK, list)
	}
}

func locationsFetch() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ce := context.Get(r, "extras")
		retval := SkeletonMessage{Message: "This is definitely a location"}
		ce.(Extra).render.JSON(w, http.StatusOK, retval)
	}
}

func locationsUpdate() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ce := context.Get(r, "extras")
		retval := SkeletonMessage{Message: "Yup, it's been updated"}
		ce.(Extra).render.JSON(w, http.StatusOK, retval)
	}
}

func locationsDelete() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ce := context.Get(r, "extras")
		retval := SkeletonMessage{Message: "Yup, it's been deleted"}
		ce.(Extra).render.JSON(w, http.StatusOK, retval)
	}
}
