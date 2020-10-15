package lib

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Server handles managing the command which runs the underlying project's executable (ie. main module program).

// StartServer - starts the Command server
func StartServer() (*exec.Cmd, error) {
	cwd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	// todo: get executable command from config, or use default executable/cwd

	mainExecutable := filepath.Base(cwd)
	//fmt.Println("Starting... " + mainExecutable)
	cmd := exec.Command("./" + mainExecutable)
	cmd.Stdout = os.Stdout
	cmd.Stderr = &ErrorWriter{}
	err = cmd.Start()

	if err != nil {
		fmt.Printf(ColorError+"Error starting command server: %s\n"+ColorReset, err)
		return nil, err
	} else {
		fmt.Print(ColorGreen + "Starting..." + ColorReset + "\n")
	}

	return cmd, nil
}

// StopServer - stops the Command server
func StopServer(server *exec.Cmd) error {
	//fmt.Println("Stopping.")
	// Server already stopped?
	if server == nil || server.Process == nil {
		return nil
	}

	// Send interrupt signal
	// TODO: Sernding interrupt on windows is not implmeneted.
	err := server.Process.Signal(os.Interrupt)

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
	return server.Process.Kill()
}
