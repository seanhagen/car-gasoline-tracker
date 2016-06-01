package main

import (
	"database/sql"
	"github.com/gchaincl/dotsql"
	"github.com/gorilla/context"
	"github.com/justinas/alice"
	"github.com/streadway/amqp"
	"github.com/unrolled/render"
	"log"
	"net/http"
)

type Extra struct {
	render  *render.Render
	db      *sql.DB
	dot     *dotsql.DotSql
	channel *amqp.Channel
	queue   amqp.Queue
}

func makeSetContextMiddleware(ch *amqp.Channel, q amqp.Queue) alice.Constructor {
	// get database connection and queries
	db, dot := getDbQuery()

	render := render.New(render.Options{
		IndentJSON: true,
	})

	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			e := Extra{
				render:  render,
				db:      db,
				dot:     dot,
				channel: ch,
				queue:   q,
			}
			context.Set(r, "extras", e)
			handler.ServeHTTP(w, r)
		})
	}
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
