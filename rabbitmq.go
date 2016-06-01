package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func getMQConn() *amqp.Connection {
	connString := getServiceURI("rabbitmq")
	if len(connString) == 0 {
		connString = *rabbitmqConnectionString
	}

	conn, err := amqp.Dial(connString)

	failOnError(err, "Failed to connect to RabbitMQ")

	return conn
}
