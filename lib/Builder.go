package lib

import (
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/rw3iss/gow/lib/utils"
)

type Builder struct {
	app         *Application
	errorWriter *utils.FormatWriter
}

func NewBuilder(app *Application) *Builder {
	return &Builder{
		app:         app,
		errorWriter: utils.NewFormatWriter(utils.ColorError + "Error:\n%s" + utils.ColorReset),
	}
}

// Build - Executes the 'go build' command and records time.
func (b *Builder) Build() error {
	start := time.Now()

	cmd := exec.Command("go", "build")
	cmd.Stdout = os.Stdout
	cmd.Stderr = b.errorWriter

	e := cmd.Run()

	if e != nil {
		return e
	}

	duration := time.Since(start)
	ms := strconv.Itoa(int(duration.Nanoseconds() / int64(1000000)))

	utils.Log("Built in " + ms + " ms. ")

	return nil
}
