package main

import (
	"fmt"

	"github.com/rw3iss/gow/lib"
)

func main() {
	run()
}

func run() {
	fmt.Println("\nInitial build...")
	err := lib.DoBuild()

	if err != nil {
		// todo: read config, if should attempt restart on error...
		fmt.Println(lib.Red + "Build error... starting watcher anyway..." + lib.Reset)
	}

	lib.StartWatcher()
}
