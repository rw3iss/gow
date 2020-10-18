package lib

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/rw3iss/gow/lib/utils"
)

// Runner handles managing the command which runs the underlying project's executable (ie. main module program).

type Runner struct {
	app *Application
	cmd *exec.Cmd // underlying server context
}

func NewRunner(app *Application) *Runner {
	return &Runner{
		app: app,
	}
}

// Start - starts the Command server
func (r *Runner) Start() error {
	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	// Todo: get executable command from config, or use default executable/cwd
	var _ = r.app.Config.Get("command", "")

	mainExecutable := filepath.Base(cwd)
	//fmt.Println("Starting... " + mainExecutable)
	cmd := exec.Command("./" + mainExecutable)
	cmd.Stdout = os.Stdout
	cmd.Stderr = NewWriter(utils.ColorError + "Error:\n%s" + utils.ColorReset)

	//&Writer{}

	// 	&struct{
	// 	Write *func
	// }{
	// 	Write: func() {},
	// }

	err = cmd.Start()

	if err != nil {
		fmt.Printf(utils.ColorError+"Error starting command server: %s\n"+utils.ColorReset, err)
		return err
	} else {
		fmt.Print(utils.ColorGreen + "Starting..." + utils.ColorReset + "\n")
	}

	r.cmd = cmd

	return nil
}

// Stop - stops the Runner
func (r *Runner) Stop() error {
	//fmt.Println("Stopping.")
	// Server already stopped?
	if r.cmd == nil || r.cmd.Process == nil {
		return nil
	}

	// Send interrupt signal
	// TODO: Sernding interrupt on windows is not implmeneted.
	err := r.cmd.Process.Signal(os.Interrupt)

	// todo: listen for external interrupts (ie. killed process)

	if err != nil {
		return err
	}

	//time.Sleep(50 * time.Millisecond)

	// Send kill signal
	status := r.cmd.Process.Kill()
	r.cmd = nil

	return status
}