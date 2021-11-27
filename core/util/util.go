package util

import (
	"fmt"
	"io"
)

// Panics when err != nil
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Panics when err != nil, and also closes the closer
func PanicOnErrCloser(err error, closer io.Closer) {
	if err != nil {
		closer.Close()
	}
	PanicOnErr(err)
}

// Logs "WARNING:" when err != nil, followed by the value of err
func WarnOnErr(err error) {
	if err != nil {
		fmt.Println("WARNING:", err)
	}
}
