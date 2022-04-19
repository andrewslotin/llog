package llog_test

import (
	"errors"
	"log"
	"os"

	"github.com/andrewslotin/llog"
)

// This example demonstrates the use of llog.NewWriter() to filter out log messages
// based on their importance levels.
func Example() {
	// configure standard logger to write warning messages and above to STDERR
	log.SetOutput(llog.NewWriter(os.Stdout, llog.WarnLevel))

	// error level, will be printed out
	log.Println("error reading file")
	// warn level, will be printed out
	log.Printf("warn: failed to close the output file: %s", errors.New("no such file"))
	// info level, will be ignored
	log.Println("something just happened")
}
