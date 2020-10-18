package lib

import (
	"fmt"
	"io"
)

type StatusWriter struct {
	formatString string
}

func (w *StatusWriter) Write(b []byte) (int, error) {
	fmt.Printf(w.formatString, b)
	return 0, nil
}

func NewWriter(formatStr string) io.Writer {
	return &StatusWriter{
		formatString: formatStr,
	}
}
