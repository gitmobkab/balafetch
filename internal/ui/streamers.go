package ui

import (
	"fmt"
	"os"
	"strings"
)

// NewInlineStreamer returns a streamer function that prints the given string on the same line, 
// overwriting the previous string. 
// This can be used to show progress updates without cluttering the terminal.
func NewInlineStreamer() func(string){
	lastPrintLength := 0

	return func(s string) {
		cleanUp := strings.Repeat(" ", lastPrintLength)
		fmt.Printf("\r%s\r%s", cleanUp, s)
		os.Stdout.Sync()
		lastPrintLength = len(s)
	}
}

// NoOpStreamer returns a streamer function that does nothing. This can be used in tests to disable streaming output.
func NoOpStreamer() func(string){
	return func(s string) {
		// do nothing
	}
}