package utils

import (
	"io"
)

type FormatWriter struct {
	formatString string
}

func (w *FormatWriter) Write(b []byte) (int, error) {
	Log(w.formatString, b)
	return 0, nil
}

func NewFormatWriter(formatStr string) io.Writer {
	return &FormatWriter{
		formatString: formatStr,
	}
}
