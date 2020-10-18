package lib

import (
	"fmt"

	"github.com/rw3iss/gow/lib/utils"
)

// Application just wraps all the parts, central bus.
type Application struct {
	Config  *Config
	Builder *Builder
	Server  *Server
	Watcher *Watcher
}

func NewApplication() *Application {
	app := &Application{}
	app.Config, _ = NewConfig("config.json")
	app.Builder = NewBuilder(app)
	app.Server = NewServer(app)
	app.Watcher = NewWatcher(app)
	return app
}

func (app *Application) Start() {
	fmt.Print(utils.ColorNotice + "\nInitial build... " + utils.ColorReset)
	err := app.Builder.Build()

	if err != nil {
		// continue to watch even with build errors...
	}

	app.Watcher.Start()
}

func (app *Application) Restart() {
	app.Server.Stop()

	fmt.Print(utils.ColorYellow + "Rebuilding (" + changedFilename + ") ... " + utils.ColorReset)
	app.Builder.Build()

	app.Server.Start()
}
