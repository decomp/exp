// Package elf provides access to ELF binary executables.
package elf

import (
	"debug/elf"
	"io"
	"io/ioutil"
	"os"
	"sort"

	"github.com/decomp/exp/bin"
	"github.com/pkg/errors"
)

// Register ELF format.
func init() {
	// Executable and Linkable Format (ELF)
	//
	//    7F 45 4C 46  |.ELF|
	const magic = "\x7FELF"
	bin.RegisterFormat("elf", magic, Parse)
}

// ParseFile parses the given ELF binary executable, reading from path.
func ParseFile(path string) (*bin.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer f.Close()
	return Parse(f)
}

// Parse parses the given ELF binary executable, reading from r.
//
// Users are responsible for closing r.
func Parse(r io.ReaderAt) (*bin.File, error) {
	// Open ELF file.
	f, err := elf.NewFile(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Parse machine architecture.
	file := &bin.File{}
	switch f.Machine {
	case elf.EM_386:
		file.Arch = bin.ArchX86_32
	case elf.EM_X86_64:
		file.Arch = bin.ArchX86_64
	}

	// Parse entry address.
	file.Entry = bin.Address(f.Entry)

	// Parse sections.
	for _, s := range f.Sections {
		data, err := s.Data()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if len(data) == 0 {
			continue
		}
		sect := &bin.Section{
			Name: s.Name,
			Addr: bin.Address(s.Addr),
			Data: data,
		}
		file.Sections = append(file.Sections, sect)
	}
	less := func(i, j int) bool {
		if file.Sections[i].Addr == file.Sections[j].Addr {
			return file.Sections[i].Name < file.Sections[j].Name
		}
		return file.Sections[i].Addr < file.Sections[j].Addr
	}
	sort.Slice(file.Sections, less)

	// Parse segments.
	for _, prog := range f.Progs {
		if prog.Type != elf.PT_LOAD {
			continue
		}
		r := prog.Open()
		data, err := ioutil.ReadAll(r)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		perm := parsePerm(prog.Flags)
		seg := &bin.Segment{
			Addr: bin.Address(prog.Vaddr),
			Data: data,
			Perm: perm,
		}
		file.Segments = append(file.Segments, seg)
	}
	less = func(i, j int) bool {
		return file.Segments[i].Addr < file.Segments[j].Addr
	}
	sort.Slice(file.Segments, less)

	// Fix section permissions.
	if len(file.Segments) > 0 {
		for _, sect := range file.Sections {
			for _, seg := range file.Segments {
				end := seg.Addr + bin.Address(len(seg.Data))
				if seg.Addr <= sect.Addr && sect.Addr < end {
					if sect.Perm == 0 {
						sect.Perm = seg.Perm
					}
				}
			}
		}
	}

	// TODO: Parse imports.

	return file, nil
}

// parsePerm returns the memory access permissions represented by the given ELF
// access permission flags.
func parsePerm(flags elf.ProgFlag) bin.Perm {
	var perm bin.Perm
	if flags&elf.PF_R != 0 {
		perm |= bin.PermR
	}
	if flags&elf.PF_W != 0 {
		perm |= bin.PermW
	}
	if flags&elf.PF_X != 0 {
		perm |= bin.PermX
	}
	return perm
}
