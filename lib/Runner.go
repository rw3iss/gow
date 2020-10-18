package lib

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/rw3iss/gow/lib/utils"
)

// Runner handles managing the command which runs the underlying project's executable (ie. the main module program).
type Runner struct {
	app         *Application
	cmd         *exec.Cmd // underlying server context
	runCommand  []interface{}
	errorWriter *utils.FormatWriter
}

func NewRunner(app *Application) *Runner {
	// use current working directory name as executable module name
	if cwd, err := os.Getwd(); err == nil {
		return &Runner{
			app:         app,
			runCommand:  app.Config.GetArray("runCommand", "./"+filepath.Base(cwd)),
			errorWriter: utils.NewFormatWriter(utils.ColorError + "Error:\n%s" + utils.ColorReset),
		}
	}

	return nil
}

// Start starts the Runner
func (r *Runner) Start() error {
	cmd := exec.Command(string(r.runCommand[0].(string)), string(r.runCommand[1].(string)))
	cmd.Stdout = os.Stdout
	cmd.Stderr = r.errorWriter

	err := cmd.Start()

	if err != nil {
		utils.Log(utils.ColorError+"Error running: %s\n"+utils.ColorReset, err)
		return err
	} else {
		utils.Log(utils.ColorGreen + "Starting..." + utils.ColorReset + "\n")
	}

	r.cmd = cmd

	return nil
}

// Stop stops the Runner
func (r *Runner) Stop() error {
	// Server already stopped?
	if r.cmd == nil || r.cmd.Process == nil {
		return nil
	}

	// Send interrupt signal
	// TODO: Sending interrupt on windows is not implmeneted.
	err := r.cmd.Process.Signal(os.Interrupt)

	// todo: listen for external interrupts (ie. killed process from somewhere else)

	if err != nil {
		return err
	}

	//time.Sleep(50 * time.Millisecond)

	// Send kill signal
	status := r.cmd.Process.Kill()
	r.cmd = nil

	return status
}
