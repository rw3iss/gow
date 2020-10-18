package lib

import (
	"io"

	"github.com/rw3iss/gow/lib/utils"
)

type StatusWriter struct {
	formatString string
}

func (w *StatusWriter) Write(b []byte) (int, error) {
	utils.Log(w.formatString, b)
	return 0, nil
}

func NewWriter(formatStr string) io.Writer {
	return &StatusWriter{
		formatString: formatStr,
	}
}
