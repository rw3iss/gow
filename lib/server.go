package lib

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/rw3iss/gow/lib/utils"
)

// Server handles managing the command which runs the underlying project's executable (ie. main module program).

type Server struct {
	app *Application
	cmd *exec.Cmd // underlying server context
}

func NewServer(app *Application) *Server {
	return &Server{
		app: app,
	}
}

// Start - starts the Command server
func (s *Server) Start() error {
	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	// Todo: get executable command from config, or use default executable/cwd
	var _ = s.app.Config.Get("command", "")

	mainExecutable := filepath.Base(cwd)
	//fmt.Println("Starting... " + mainExecutable)
	cmd := exec.Command("./" + mainExecutable)
	cmd.Stdout = os.Stdout
	cmd.Stderr = &ErrorWriter{}
	err = cmd.Start()

	if err != nil {
		fmt.Printf(utils.ColorError+"Error starting command server: %s\n"+utils.ColorReset, err)
		return err
	} else {
		fmt.Print(utils.ColorGreen + "Starting..." + utils.ColorReset + "\n")
	}

	s.cmd = cmd

	return nil
}

// Stop - stops the Command server
func (s *Server) Stop() error {
	//fmt.Println("Stopping.")
	// Server already stopped?
	if s.cmd == nil || s.cmd.Process == nil {
		return nil
	}

	// Send interrupt signal
	// TODO: Sernding interrupt on windows is not implmeneted.
	err := s.cmd.Process.Signal(os.Interrupt)

	// todo: listen for external interrupts (ie. killed process)

	if err != nil {
		return err
	}

	// state, e := server.Process.Wait()
	// if e != nil {
	// 	fmt.Printf("error waiting %s", e)
	// 	return e
	// }

	// fmt.Printf("State %v", state)

	//time.Sleep(50 * time.Millisecond)

	// Send kill signal
	r := s.cmd.Process.Kill()
	s.cmd = nil

	return r
}
