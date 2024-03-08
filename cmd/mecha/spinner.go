package main

import (
	"io"

	"github.com/briandowns/spinner"
)

type spinWriter struct {
	s       *spinner.Spinner
	out     io.Writer
	written bool
}

func (sw *spinWriter) Write(p []byte) (n int, err error) {
	if !sw.written {
		sw.s.Stop()
	}

	n, err = sw.out.Write(p)
	return
}
