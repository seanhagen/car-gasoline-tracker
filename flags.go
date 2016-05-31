package main

import (
	"flag"
)

var (
	// http stuff
	serverPortFlag    = flag.Int("server-port", 8080, "Port to run web server on")
	serverAddressFlag = flag.String("server-addr", "127.0.0.1", "Address to listen on")

	// database stuff
	databaseConnectionString = flag.String(
		"db-conn-string",
		"postgres://gasman:gasman_test@localhost/gasman_dev",
		"Database connection string",
	)

	googleMapsApiKey = flag.String(
		"google-maps-api-key",
		"key-goes-here",
		"Google Maps API key",
	)
)
