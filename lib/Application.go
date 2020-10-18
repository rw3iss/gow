package lib

import (
	"github.com/rw3iss/gow/lib/utils"
)

// Application just wraps all the parts, central bus.
type Application struct {
	Config  *Config
	Builder *Builder
	Runner  *Runner
	Watcher *Watcher
}

func NewApplication() *Application {
	app := &Application{}
	app.Config, _ = NewConfig("config.json")
	app.Builder = NewBuilder(app)
	app.Runner = NewRunner(app)
	app.Watcher = NewWatcher(app)
	return app
}

func (app *Application) Start() {
	utils.Log(utils.ColorNotice + "\nInitial build... " + utils.ColorReset)
	err := app.Builder.Build()

	if err != nil {
		// continue to watch even with build errors...
	}

	app.Watcher.Start()
}

func (app *Application) Restart() {
	app.Runner.Stop()

	utils.Log(utils.ColorYellow + "Rebuilding (" + changedFilename + ") ... " + utils.ColorReset)
	app.Builder.Build()

	app.Runner.Start()
}
