package lib

import (
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/rw3iss/gow/lib/utils"
)

type Watcher struct {
	app *Application
}

func NewWatcher(app *Application) *Watcher {
	return &Watcher{
		app: app,
	}
}

// Start - Starts the watching of the diretory and notifies iof changes
func (w *Watcher) Start() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		utils.Log(err.Error())
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
					w.app.Restart(event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				utils.Log("Watcher error: %s", err.Error())
			}
		}
	}()

	// watch the current directory the command was run in
	// Todo: read watch dir from Config
	watchDir := "./"
	err = recurseWatchDirs(watcher, watchDir)
	if err != nil {
		return
	}

	utils.Log("\n"+utils.ColorNotice+"Watching: %s"+utils.ColorReset+"\n", watchDir)

	// start the target executable
	w.app.Runner.Start()

	<-done
}

func recurseWatchDirs(watcher *fsnotify.Watcher, dir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				utils.Log("Error reading file: %s", err.Error())
				return err
			}
			if info.IsDir() {
				watcher.Add(path)
			}
			return nil
		})

	if err != nil {
		utils.Log(err.Error())
		return err
	}

	return nil
}
