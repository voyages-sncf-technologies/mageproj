package util

import (
	"os"
	"strconv"
)

// Verbose reports whether a magefile was run with the verbose flag.
func Verbose() bool {
	b, _ := strconv.ParseBool(os.Getenv("MAGEFILE_VERBOSE"))
	return b
}

// Debug reports whether a magefile was run with the debug flag.
func Debug() bool {
	b, _ := strconv.ParseBool(os.Getenv("MAGEFILE_DEBUG"))
	return b
}

// GoCmd reports the command to use to build go code. By default it is
// the "go" binary in the PATH.
func GoCmd() string {
	if cmd := os.Getenv("MAGEFILE_GOCMD"); cmd != "" {
		return cmd
	}
	return "go"
}

// GitCmd reports the command to use to extract git info. By default it is
// the "go" binary in the PATH.
func GitCmd() string {
	return "git"
}