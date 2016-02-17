package main

import (
	"database/sql"
	"github.com/seanhagen/gas-web/Godeps/_workspace/src/github.com/gchaincl/dotsql"
	"github.com/seanhagen/gas-web/Godeps/_workspace/src/github.com/gorilla/context"
	"github.com/seanhagen/gas-web/Godeps/_workspace/src/github.com/justinas/alice"
	"github.com/seanhagen/gas-web/Godeps/_workspace/src/github.com/unrolled/render"
	"log"
	"net/http"
)

type Extra struct {
	render *render.Render
	db     *sql.DB
	dot    *dotsql.DotSql
}

func makeSetContextMiddleware(ren *render.Render, db *sql.DB, dot *dotsql.DotSql) alice.Constructor {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			e := Extra{
				render: ren,
				db:     db,
				dot:    dot,
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
