package lib

import (
	"fmt"
)

type StatusWriter struct {
	FormatString string
}

func (w *StatusWriter) Write(b []byte) (int, error) {
	fmt.Printf(w.FormatString, b)
	return 0, nil
}

func NewWriter(formatStr string) interface{ Write(b []byte) (int, error) } {
	return &StatusWriter{
		FormatString: formatStr,
	}
}
