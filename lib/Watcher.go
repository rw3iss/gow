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

// Start starts the watching of the diretory and notifies iof changes
func (w *Watcher) Start() {
	fileWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		utils.Log(err.Error())
		return
	}
	defer fileWatcher.Close()

	done := make(chan bool)

	// coroutine to respond to file events:
	go func() {
		for {
			select {
			case event, ok := <-fileWatcher.Events:
				if !ok {
					return
				}
				//log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					w.app.Restart(event.Name)
				}
			case err, ok := <-fileWatcher.Errors:
				if !ok {
					return
				}
				utils.Log("Watcher error: %s", err.Error())
			}
		}
	}()

	// watch the target or current directory:
	watchDir := w.app.Config.Get("watchDir", "./")

	if _, err := os.Stat(watchDir); os.IsNotExist(err) {
		panic("\nConfigured watchDir directory does not exist: " + watchDir)
	}

	err = w._recurseWatchDirs(fileWatcher, watchDir)
	if err != nil {
		return
	}

	utils.Log("\n"+utils.ColorNotice+"Watching: %s"+utils.ColorReset+"\n", watchDir)

	// start the target executable
	w.app.Start()

	<-done
}

// helpers to add all subdirectories to the watcher
func (w *Watcher) _recurseWatchDirs(fileWatcher *fsnotify.Watcher, dir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				utils.Log("Error reading file: %s", err.Error())
				return err
			}
			if info.IsDir() {
				fileWatcher.Add(path)
			}
			return nil
		})

	if err != nil {
		utils.Log(err.Error())
		return err
	}

	return nil
}
