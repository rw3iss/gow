package lib

import "fmt"

type ErrorWriter struct {
}

func (cw *ErrorWriter) Write(b []byte) (int, error) {
	fmt.Printf(ColorError+"Build error:\n%s"+ColorReset, b)
	return 0, nil
}
