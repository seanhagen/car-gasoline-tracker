package main

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func recordsList() httprouter.Handle {
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

func recordsFetch() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ce := context.Get(r, "extras")
		retval := SkeletonMessage{Message: "This is definitely a record"}
		ce.(Extra).render.JSON(w, http.StatusOK, retval)
	}
}

func recordsCreate() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ce := context.Get(r, "extras")
		msg := SkeletonMessage{Message: "Yup, created!"}
		ce.(Extra).render.JSON(w, http.StatusOK, msg)
	}
}

func recordsUpdate() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ce := context.Get(r, "extras")
		retval := SkeletonMessage{Message: "Yup, it's been updated"}
		ce.(Extra).render.JSON(w, http.StatusOK, retval)
	}
}

func recordtsDelete() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ce := context.Get(r, "extras")
		retval := SkeletonMessage{Message: "Yup, it's been deleted"}
		ce.(Extra).render.JSON(w, http.StatusOK, retval)
	}
}
