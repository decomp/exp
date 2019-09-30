// Note, the file format registration implementation of this package is heavily
// inspired by the image package of the Go standard library, which is governed
// by a BSD license.

package bin

import (
	"io"
	"log"
	"os"

	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
)

// Loggers.
var (
	// dbg is a logger with the "bin:" prefix which logs debug messages to
	// standard error.
	dbg = log.New(os.Stderr, term.MagentaBold("bin:")+" ", 0)
	// warn is a logger with the "bin:" prefix which logs warning messages to
	// standard error.
	warn = log.New(os.Stderr, term.RedBold("bin:")+" ", 0)
)

// RegisterFormat registers a binary executable format for use by Parse. Name is
// the name of the format, like "pe" or "elf". Magic is the magic prefix that
// identifies the format's encoding. The magic string can contain "?" wildcards
// that each match any one byte.
func RegisterFormat(name, magic string, parse func(io.ReaderAt) (*File, error)) {
	formats = append(formats, format{name: name, magic: magic, parse: parse})
}

// formats is the list of registered formats.
var formats []format

// A format holds a binary executable format's name, magic header and how to
// parse it.
type format struct {
	// Name of the binary executable format.
	name string
	// Magic prefix that identifies the format's encoding. The magic string can
	// contain "?" wildcards that each match any one byte.
	magic string
	// parse parses the given binary executable, reading from r.
	parse func(r io.ReaderAt) (*File, error)
}

// ParseFile parses the given binary executable, reading from path.
func ParseFile(path string) (*File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer f.Close()
	return Parse(f)
}

// Parse parses the given binary executable, reading from r.
//
// Users are responsible for closing r.
func Parse(r io.ReaderAt) (*File, error) {
	for _, format := range formats {
		buf := make([]byte, len(format.magic))
		if n, err := r.ReadAt(buf, 0); err != nil {
			if errors.Cause(err) == io.EOF && n < len(format.magic) {
				warn.Printf("skip %q format (read %d of %d bytes required for magic identification)", format.name, n, len(format.magic))
				continue
			}
			return nil, errors.WithStack(err)
		}
		if match(format.magic, buf) {
			return format.parse(r)
		}
	}
	return nil, errors.New("unknown binary executable format;\n\ttip: try loading as raw binary executable")
}

// match reports whether magic matches b. The magic string can contain "?"
// wildcards that each match any one byte.
func match(magic string, buf []byte) bool {
	if len(magic) != len(buf) {
		return false
	}
	for i, b := range buf {
		if magic[i] != '?' && magic[i] != b {
			return false
		}
	}
	return true
}
