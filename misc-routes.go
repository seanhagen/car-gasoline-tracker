package main

import (
	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/gorilla/context"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/streadway/amqp"
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

func appenvRoute() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		appEnv, _ := cfenv.Current()
		ce := context.Get(r, "extras")
		if appEnv != nil {
			ce.(Extra).render.JSON(w, http.StatusOK, appEnv.Services)
		} else {
			msg := SkeletonMessage{Message: "Blergh: got no appenv."}
			ce.(Extra).render.JSON(w, http.StatusOK, msg)
		}
	}
}

func postToQueue() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ce := context.Get(r, "extras")
		err := ce.(Extra).channel.Publish(
			"",
			ce.(Extra).queue.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("hello"),
			},
		)
		if err != nil {
			panic(err)
		}

		msg := SkeletonMessage{Message: "Delivered message to queue!"}
		ce.(Extra).render.JSON(w, http.StatusOK, msg)

	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func websocketRoute() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ce := context.Get(r, "extras")
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			panic(err)
		}

		msgs, err := ce.(Extra).channel.Consume(
			ce.(Extra).queue.Name,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			panic(err)
		}

		for {
			for d := range msgs {
				conn.WriteMessage(websocket.TextMessage, d.Body)
			}
		}
	}
}
