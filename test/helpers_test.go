package test

import (
	"bytes"
	"io"
	"os"
)

// captureOutput captures what gets printed to stdout
// Takes a function that prints, returns the captured string
func captureOutput(f func()) string {
	// Save original stdout
	oldStdout := os.Stdout

	// Create a pipe (r = read end, w = write end)
	r, w, _ := os.Pipe()

	// Redirect stdout to the write end of the pipe
	os.Stdout = w

	// Run the function (it will print to our pipe instead of terminal)
	f()

	// Close the write end and restore original stdout
	w.Close()
	os.Stdout = oldStdout

	// Read everything from the pipe into a buffer
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Return the captured output as a string
	return buf.String()
}
