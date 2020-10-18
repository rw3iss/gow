package utils

type FormatWriter struct {
	formatString string
}

func (w *FormatWriter) Write(b []byte) (int, error) {
	Log(w.formatString, b)
	return 0, nil
}

func NewFormatWriter(formatStr string) *FormatWriter {
	return &FormatWriter{
		formatString: formatStr,
	}
}
