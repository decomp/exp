// Package bin provides a uniform representation of binary executables.
package bin

import (
	"log"
	"os"

	"github.com/mewkiz/pkg/term"
)

var (
	// dbg is a logger with the "bin:" prefix which logs debug messages to
	// standard error.
	dbg = log.New(os.Stderr, term.MagentaBold("bin:")+" ", 0)
	// warn is a logger with the "bin:" prefix which logs warning messages to
	// standard error.
	warn = log.New(os.Stderr, term.RedBold("bin:")+" ", 0)
)
