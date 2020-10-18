package utils

import "fmt"

func Debug(msg string, data ...interface{}) {
	// todo: check config if debug is enabled...
	fmt.Printf(msg, data...)
}

func Log(msg string, data ...interface{}) {
	fmt.Printf(msg, data...)
}
