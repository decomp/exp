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
// The entry point and base address are both 0 by default. To specify a custom
// entry point, set file.Entry, and to specify a custom base address, set
// file.Sections[0].Addr.
func ParseFile(path string, arch bin.Arch) (*bin.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer f.Close()
	return Parse(f, arch)
}

// Parse parses the given raw binary executable, reading from r.
//
// The entry point and base address are both 0 by default. To specify a custom
// entry point, set file.Entry, and to specify a custom base address, set
// file.Sections[0].Addr.
func Parse(r io.Reader, arch bin.Arch) (*bin.File, error) {
	// Parse segments.
	file := &bin.File{
		Arch: arch,
	}
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	seg := &bin.Section{
		Addr:     0,
		Data:     data,
		FileSize: len(data),
		MemSize:  len(data),
		Perm:     bin.PermR | bin.PermW | bin.PermX,
	}
	file.Sections = append(file.Sections, seg)
	return file, nil
}
