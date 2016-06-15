package main

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"net/http"
)

func locationsIndex() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ren := context.Get(r, "render").(*render.Render)
		var list []SkeletonMessage = []SkeletonMessage{
			SkeletonMessage{Message: "This is a thing"},
			SkeletonMessage{Message: "This is a thing"},
			SkeletonMessage{Message: "This is a thing"},
			SkeletonMessage{Message: "This is a thing"},
			SkeletonMessage{Message: "This is a thing"},
			SkeletonMessage{Message: "This is a thing"},
			SkeletonMessage{Message: "This is a thing"},
		}
		ren.JSON(w, http.StatusOK, list)
	}
}

func locationsGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ren := context.Get(r, "render").(*render.Render)
		retval := SkeletonMessage{Message: "This is definitely a location"}
		ren.JSON(w, http.StatusOK, retval)
	}
}

func locationsCreate() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ren := context.Get(r, "render").(*render.Render)
		msg := SkeletonMessage{Message: "Yup, created!"}
		ren.JSON(w, http.StatusOK, msg)
	}
}

func locationsUpdate() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ren := context.Get(r, "render").(*render.Render)
		retval := SkeletonMessage{Message: "Yup, it's been updated"}
		ren.JSON(w, http.StatusOK, retval)
	}
}

func locationsDelete() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ren := context.Get(r, "render").(*render.Render)
		retval := SkeletonMessage{Message: "Yup, it's been deleted"}
		ren.JSON(w, http.StatusOK, retval)
	}
}
