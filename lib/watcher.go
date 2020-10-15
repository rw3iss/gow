package lib

import (
	"fmt"
	"os/exec"

	"github.com/fsnotify/fsnotify"
)

var server *exec.Cmd
var changedFilename string

func restartApp() {
	fmt.Print("Restarting (" + changedFilename + ") ...")

	// Stop (we can stop in a separate thread)
	go StopServer(server)

	// Rebuild
	_ = DoBuild()

	// Restart
	server, _ = StartServer()
}

// StartWatcher - Starts the watching of the diretory and notifies iof changes
func StartWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer watcher.Close()

	done := make(chan bool)

	// coroutine to respond to file events:
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				//log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					changedFilename = event.Name
					go restartApp()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Printf("Watcher error: %s", err.Error())
			}
		}
	}()

	// start the command server
	server, err = StartServer()
	if err != nil {
		fmt.Print(err)
		return
	}

	// watch the current directory the command was run in
	watchDir := "./"
	err = watcher.Add(watchDir)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	fmt.Printf("Watching: %s\n\n", watchDir)

	<-done
}
