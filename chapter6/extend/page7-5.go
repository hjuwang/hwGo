package main

import (
	"bytes"
	"io"
)

const debug = false

func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		// ...use buf...
	}
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ...do something...
	if out != nil {
		_, err := out.Write([]byte("done!\n"))
		if err != nil {
			return
		}
	}
}
