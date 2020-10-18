package lib

import (
	"fmt"

	"github.com/rw3iss/gow/lib/utils"
)

// ErrorWriter exposes an interface to write a stream of bytes to the console/stdout, formatted as an Error.
// Todo: generacize it as StatusWriter and pass status flag/enum to handle different coloring in one location.
type ErrorWriter struct {
}

func (cw *ErrorWriter) Write(b []byte) (int, error) {
	fmt.Printf(utils.ColorError+"Build error:\n%s"+utils.ColorReset, b)
	return 0, nil
}
