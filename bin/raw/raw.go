// Package raw provides access to raw binary executables.
package raw

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/decomp/exp/bin"
	"github.com/pkg/errors"
)

// ParseFile parses the given raw binary executable, reading from path.
//
// Users are responsible for specifying file.Arch. By default the entry point
// and base address are both 0. To specify a custom entry point, set file.Entry,
// and to specify a custom base address, set file.Segments[0].Addr.
func ParseFile(path string) (*bin.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer f.Close()
	return Parse(f)
}

// Parse parses the given raw binary executable, reading from r.
//
// Users are responsible for specifying file.Arch. By default the entry point
// and base address are both 0. To specify a custom entry point, set file.Entry,
// and to specify a custom base address, set file.Segments[0].Addr.
func Parse(r io.Reader) (*bin.File, error) {
	// Parse segments.
	file := &bin.File{}
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	seg := &bin.Section{
		Addr: 0,
		Data: data,
		Perm: bin.PermR | bin.PermW | bin.PermX,
	}
	file.Segments = append(file.Segments, seg)
	return file, nil
}
