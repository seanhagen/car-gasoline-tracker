package main

import (
	"github.com/gorilla/context"
	"github.com/justinas/alice"
	"github.com/unrolled/render"
	"log"
	"net/http"
)

func makeContextMiddleware() alice.Constructor {
	render := render.New(render.Options{
		IndentJSON: true,
	})

	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			context.Set(r, "render", render)
			handler.ServeHTTP(w, r)
		})
	}
}

func logMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
