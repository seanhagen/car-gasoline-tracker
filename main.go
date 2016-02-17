package main

import (
	"github.com/seanhagen/gas-web/Godeps/_workspace/src/github.com/vharitonsky/iniflags"
)

func main() {
	// set up recovery
	defer recoverPanic()

	// intialize flags
	iniflags.Parse()

	// start web server
	server()
}
