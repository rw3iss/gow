package lib

import (
	"fmt"
)

type StatusWriter struct {
	formatString string
}

func (w *StatusWriter) Write(b []byte) (int, error) {
	fmt.Printf(w.formatString, b)
	return 0, nil
}

func NewWriter(formatStr string) interface{ Write(b []byte) (int, error) } {
	return &StatusWriter{
		formatString: formatStr,
	}
}
