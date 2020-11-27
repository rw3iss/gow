package lib

import (
	"log"
	"os"

	"github.com/rw3iss/gow/lib/utils"
)

// Application just wraps all the parts, central bus.
type Application struct {
	Config  *utils.Config
	Builder *Builder
	Runner  *Runner
	Watcher *Watcher
}

func NewApplication() *Application {
	app := &Application{}

	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Could not get working directory: %v", err)
	}
	log.Printf("Log path: %s", path)

	app.Config, _ = utils.NewConfig(path + "/config.json")
	app.Builder = NewBuilder(app)
	app.Runner = NewRunner(app)
	app.Watcher = NewWatcher(app)
	return app
}

// Init does the initial build, and starts the file watcher.
func (app *Application) Init() {
	utils.Log(utils.ColorNotice + "\nInitial build... " + utils.ColorReset)
	err := app.Builder.Build()

	if err != nil {
		// continue to watch even with build errors...
	}

	app.Watcher.Start()
}

// Start starts the target program.
func (app *Application) Start() {
	app.Runner.Start()
}

// Retart rebuilds and restarts the target program.
func (app *Application) Restart(changedFilename string) {
	app.Runner.Stop()

	utils.Log(utils.ColorYellow + "Rebuilding (" + changedFilename + ") ... " + utils.ColorReset)
	app.Builder.Build()

	app.Runner.Start()
}
