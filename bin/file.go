// Package bin provides a uniform representation of binary executables.
package bin

import (
	"fmt"
	"sort"
)

// A File is a binary exectuable.
type File struct {
	// Machine architecture specifying the assembly instruction set.
	Arch Arch
	// Entry point of the executable.
	Entry Address
	// Sections of the exectuable.
	Sections []*Section
	// Segments of the exectuable.
	Segments []*Section
}

// Data returns the data starting at the specified address of the binary
// executable.
func (file *File) Data(addr Address) []byte {
	if len(file.Sections) > 0 {
		data, ok := locateData(addr, file.Sections)
		if ok {
			return data
		}
	}
	if len(file.Segments) > 0 {
		data, ok := locateData(addr, file.Segments)
		if ok {
			return data
		}
	}
	panic(fmt.Errorf("unable to locate data at address %v", addr))
}

// locateData tries to locate the data starting at the specified address by
// searching through the given sections. The boolean return value indicates
// success.
func locateData(addr Address, sects []*Section) ([]byte, bool) {
	// Find the first section who's end address is greater than addr.
	less := func(i int) bool {
		sect := sects[i]
		return addr < sect.Addr+Address(len(sect.Data))
	}
	index := sort.Search(len(sects), less)
	if 0 <= index && index < len(sects) {
		sect := sects[index]
		if sect.Addr <= addr && addr < sect.Addr+Address(len(sect.Data)) {
			offset := addr - sect.Addr
			return sect.Data[offset:], true
		}
	}
	return nil, false
}

// Arch represents the set of machine architectures.
type Arch uint

// Machine architectures.
const (
	// ArchX86_32 represents the 32-bit x86 machine architecture, as used by
	// Intel and AMD.
	ArchX86_32 Arch = 1 + iota
	// ArchX86_64 represents the 64-bit x86-64 machine architecture, as used by
	// Intel and AMD.
	ArchX86_64
)

// A Section represents a continuous section of memory.
type Section struct {
	// Section name; or empty if memory segments.
	Name string
	// Start address of the section.
	Addr Address
	// Section contents.
	Data []byte
	// Access permissions of the section in memory.
	Perm Perm
}

// Perm specifies the access permissions of a segment or section in memory.
type Perm uint8

// Access permissions.
const (
	// PermX specifies that the memory is executable.
	PermX Perm = 0x1
	// PermW specifies that the memory is writeable.
	PermW Perm = 0x2
	// PermR specifies that the memory is readable.
	PermR Perm = 0x4
)

// String returns the string representation of the access permissions.
func (perm Perm) String() string {
	r := "-"
	if perm&PermR != 0 {
		r = "r"
	}
	w := "-"
	if perm&PermW != 0 {
		w = "w"
	}
	x := "-"
	if perm&PermX != 0 {
		x = "x"
	}
	return r + w + x
}
