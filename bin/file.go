// Package bin provides a uniform representation of binary executables.
package bin

import (
	"fmt"
	"sort"
	"strings"
)

// A File is a binary exectuable.
type File struct {
	// Machine architecture specifying the assembly instruction set.
	Arch Arch
	// Entry point of the executable.
	Entry Address
	// Sections (and segments) of the exectuable.
	Sections []*Section
	// Function imports.
	Imports map[Address]string
	// Function exports.
	Exports map[Address]string
}

// Code returns the code starting at the specified address of the binary
// executable.
func (file *File) Code(addr Address) []byte {
	if len(file.Sections) > 0 {
		code, ok := locateCode(addr, file.Sections)
		if ok {
			return code
		}
	}
	panic(fmt.Errorf("unable to locate code at address %v", addr))
}

// locateCode tries to locate the code starting at the specified address by
// searching through the given sections. The boolean return value indicates
// success.
//
// pre-condition: sects must be sorted in ascending order.
func locateCode(addr Address, sects []*Section) ([]byte, bool) {
	// Find the first section who's end address is greater than addr.
	less := func(i int) bool {
		sect := sects[i]
		return addr < sect.Addr+Address(len(sect.Data))
	}
	index := sort.Search(len(sects), less)
	for i := index; index < len(sects); i++ {
		sect := sects[i]
		if sect.Perm&PermX == 0 {
			// skip non-executable section.
			continue
		}
		if sect.Addr <= addr && addr < sect.Addr+Address(len(sect.Data)) {
			offset := addr - sect.Addr
			return sect.Data[offset:], true
		}
	}
	return nil, false
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
	panic(fmt.Errorf("unable to locate data at address %v", addr))
}

// locateData tries to locate the data starting at the specified address by
// searching through the given sections. The boolean return value indicates
// success.
//
// pre-condition: sects must be sorted in ascending order.
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
	// ArchMIPS_32 represents the 32-bit MIPS machine architecture.
	ArchMIPS_32
	// ArchPowerPC_32 represents the 32-bit PowerPC machine architecture.
	ArchPowerPC_32
)

// BitSize returns the bit size of the machine architecture.
func (arch Arch) BitSize() int {
	m := map[Arch]int{
		// 32-bit architectures.
		ArchX86_32:     32,
		ArchMIPS_32:    32,
		ArchPowerPC_32: 32,
		// 64-bit architectures.
		ArchX86_64: 64,
	}
	if n, ok := m[arch]; ok {
		return n
	}
	panic(fmt.Errorf("support for machine architecture %v not yet implemented", uint(arch)))
}

// Set sets arch to the machine architecture represented by s.
func (arch *Arch) Set(s string) error {
	m := map[string]Arch{
		"x86_32":     ArchX86_32,
		"x86_64":     ArchX86_64,
		"MIPS_32":    ArchMIPS_32,
		"PowerPC_32": ArchPowerPC_32,
	}
	if v, ok := m[s]; ok {
		*arch = v
		return nil
	}
	var ss []string
	for s := range m {
		ss = append(ss, s)
	}
	sort.Strings(ss)
	return fmt.Errorf("support for machine architecture %q not yet implemented;\n\tsupported machine architectures: %v", s, strings.Join(ss, ", "))
}

// String returns a string representation of the machine architecture.
func (arch Arch) String() string {
	m := map[Arch]string{
		ArchX86_32:     "x86_32",
		ArchX86_64:     "x86_64",
		ArchMIPS_32:    "MIPS_32",
		ArchPowerPC_32: "PowerPC_32",
	}
	if s, ok := m[arch]; ok {
		return s
	}
	if arch == 0 {
		return "machine architecture NONE"
	}
	panic(fmt.Errorf("support for machine architecture %v not yet implemented", uint(arch)))
}

// A Section represents a continuous section of memory.
type Section struct {
	// Section name; or empty if unnamed section or memory segment.
	Name string
	// Start address of the section.
	Addr Address
	// File offset of the section.
	Offset uint64
	// Section contents.
	Data []byte
	// Size in bytes of the section contents when loaded into memory. The virtual
	// size is larger than the raw size (i.e. len(sect.Data)) for sections padded
	// to section alignment in the executable file, and smaller than the raw size
	// for sections containing uninitialized data not part of the executable
	// file.
	MemSize int
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
