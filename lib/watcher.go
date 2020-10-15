package lib

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

var server *exec.Cmd
var changedFilename string

func restartApp() {
	fmt.Print(ColorYellow + "Rebuilding (" + changedFilename + ") ... " + ColorReset)

	// Stop (we can stop in a separate thread)
	StopServer(server)

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
					restartApp()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Printf("Watcher error: %s", err.Error())
			}
		}
	}()

	// watch the current directory the command was run in
	watchDir := "./"

	err = recurseWatchDirs(watcher, watchDir)
	if err != nil {
		return
	}

	fmt.Printf("\n"+ColorNotice+"Watching: %s"+ColorReset+"\n", watchDir)

	// start the command server
	server, err = StartServer()
	if err != nil {
		fmt.Print(err)
		return
	}

	<-done
}

func recurseWatchDirs(watcher *fsnotify.Watcher, dir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("Error reading file: %s", err.Error())
				return err
			}
			if info.IsDir() {
				//fmt.Printf("Watching: %s\n", path)
				watcher.Add(path)
			}
			return nil
		})

	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	return nil
}
