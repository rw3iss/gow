package lib

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type CustomWriter struct {
}

func (cw *CustomWriter) Write(b []byte) (int, error) {
	fmt.Println("CustomWriter.Write")
	return 0, nil
}

// DoBuild - Executes the 'go build' command and records time.
func DoBuild() error {
	start := time.Now()

	cmd := exec.Command("go", "build")
	cmd.Stdout = os.Stdout

	var errorWriter = &CustomWriter{}

	cmd.Stderr = errorWriter //os.Stderr

	fmt.Println(HiRed)

	e := cmd.Run()

	fmt.Println(Reset)

	if e != nil {
		//fmt.Printf("\nError in build: %s", e)
		return e
	}

	duration := time.Since(start)
	ms := strconv.Itoa(int(duration.Nanoseconds() / int64(1000000)))

	// if err != nil {
	// 	return err
	// }

	fmt.Println("Built in " + ms + " ms.\n")

	return nil
}
