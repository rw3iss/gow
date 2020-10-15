package main

import (
	"fmt"

	"github.com/rw3iss/gow/lib"
)

func main() {
	run()
}

func run() {
	// try to read in Config values
	// var config, _ = lib.NewConfig("config.json")
	// if (config.Get()) {
	// }

	fmt.Print(lib.ColorNotice + "\nInitial build... " + lib.ColorReset)
	err := lib.DoBuild()

	if err != nil {
		// continue to watch even with build errors...
		//fmt.Println(lib.ResetColor + "Starting watcher..." + lib.ResetColor + "\n")
	}

	lib.StartWatcher()
}
