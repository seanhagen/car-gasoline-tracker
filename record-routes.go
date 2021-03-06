package main

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"net/http"
)

func recordsList() httprouter.Handle {
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

func recordsGet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ren := context.Get(r, "render").(*render.Render)
		// retval := SkeletonMessage{Message: "This is definitely a record"}
		// ren.JSON(w, http.StatusOK, retval)
		rec, _ := NewRecord(nil)
		fmt.Printf("Record: %#v\n", rec)
		rec.Create(r)
		fmt.Printf("Record now: %#v\n", rec)
		ren.JSON(w, http.StatusOK, rec)
	}
}

func recordsCreate() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ren := context.Get(r, "render").(*render.Render)
		// msg := SkeletonMessage{Message: "Yup, created!"}

		rec, _ := NewRecord(nil)
		ren.JSON(w, http.StatusOK, rec.Create(r))
	}
}

func recordsUpdate() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ren := context.Get(r, "render").(*render.Render)
		retval := SkeletonMessage{Message: "Yup, it's been updated"}
		ren.JSON(w, http.StatusOK, retval)
	}
}

func recordsDelete() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ren := context.Get(r, "render").(*render.Render)
		retval := SkeletonMessage{Message: "Yup, it's been deleted"}
		ren.JSON(w, http.StatusOK, retval)
	}
}
