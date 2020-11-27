package main

import (
	"github.com/rw3iss/gow/lib"
)

// global app context
var app *lib.Application

func main() {
	app = lib.NewApplication()
	app.Init()
}
