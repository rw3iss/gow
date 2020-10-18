package lib

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/rw3iss/gow/lib/utils"
)

type Watcher struct {
	app *Application
}

//var server *exec.Cmd
var changedFilename string

func NewWatcher(app *Application) *Watcher {
	return &Watcher{
		app: app,
	}
}

// StartWatcher - Starts the watching of the diretory and notifies iof changes
func (w *Watcher) Start() {
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
					w.app.Restart()
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

	fmt.Printf("\n"+utils.ColorNotice+"Watching: %s"+utils.ColorReset+"\n", watchDir)

	// start the target Runner
	w.app.Runner.Start()

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
