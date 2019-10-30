package bin

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
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
		if code, ok := locateCode(addr, file.Sections); ok {
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

//go:generate stringer -linecomment -type Arch
//go:generate string2enum -samepkg -linecomment -type Arch

// Arch represents the set of machine architectures.
type Arch uint16

// Machine architectures.
const (
	// ArchX86_32 represents the 32-bit x86 machine architecture, as used by
	// Intel and AMD.
	ArchX86_32 Arch = 1 + iota // x86_32
	// ArchX86_64 represents the 64-bit x86-64 machine architecture, as used by
	// Intel and AMD.
	ArchX86_64 // x86_64
	// ArchMIPS_32 represents the 32-bit MIPS machine architecture.
	ArchMIPS_32 // MIPS_32
	// ArchARM_32 represents the 32-bit ARM machine architecture.
	ArchARM_32 // ARM_32
	// ArchARM_64 represents the 64-bit ARM machine architecture.
	ArchARM_64 // ARM_64
	// ArchPowerPC_32 represents the 32-bit PowerPC machine architecture.
	ArchPowerPC_32 // PowerPC_32
	// ArchPowerPC_64BE represents the 64-bit PowerPC machine architecture
	// encoded as big endian.
	ArchPowerPC_64BE // PowerPC_64 big endian
	// ArchPowerPC_64LE represents the 64-bit PowerPC machine architecture
	// encoded as little endian.
	ArchPowerPC_64LE // PowerPC_64 little endian

	// First and last machine architectures.
	archFirst = ArchX86_32
	archLast  = ArchPowerPC_64LE
)

// bitSize maps from machine architecture to bit size.
var bitSize = map[Arch]int{
	// 32-bit architectures.
	ArchX86_32:     32,
	ArchMIPS_32:    32,
	ArchPowerPC_32: 32,
	ArchARM_32:     32,
	// 64-bit architectures.
	ArchARM_64:       64,
	ArchX86_64:       64,
	ArchPowerPC_64BE: 64,
	ArchPowerPC_64LE: 64,
}

// BitSize returns the bit size of the machine architecture.
func (arch Arch) BitSize() int {
	if n, ok := bitSize[arch]; ok {
		return n
	}
	panic(fmt.Errorf("support for machine architecture %v not yet implemented", uint16(arch)))
}

// Set sets arch to the machine architecture represented by s.
func (arch *Arch) Set(s string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			var ss []string
			for a := archFirst; a <= archLast; a++ {
				ss = append(ss, strconv.Quote(a.String()))
			}
			err = errors.Wrapf(e.(error), "valid Arch enums are: %v", strings.Join(ss, ", "))
		}
	}()
	*arch = ArchFromString(s)
	return err
}

// UnmarshalText unmarshals the text into arch.
func (arch *Arch) UnmarshalText(text []byte) error {
	return arch.Set(string(text))
}

// MarshalText returns the textual representation of arch.
func (arch Arch) MarshalText() ([]byte, error) {
	return []byte(arch.String()), nil
}

// A Section represents a continuous section of memory.
type Section struct {
	// Section name; or empty if unnamed section or memory segment.
	Name string
	// Start address of the section.
	Addr Address
	// File offset of the section.
	Offset uint64
	// Section contents of initialized data; excluding section alignment padding.
	Data []byte
	// Size in bytes of the section contents in the executable file; including
	// section alignment padding.
	//
	// FileSize is larger than MemSize for sections padded to section alignment
	// in the executable file.
	FileSize int
	// Size in bytes of the section contents when loaded into memory.
	//
	// MemSize is larger than FileSize for sections containing uninitialized data
	// not part of the executable file.
	MemSize int
	// Access permissions of the section in memory.
	Perm Perm
}

// Perm specifies the access permissions of a segment or section in memory.
type Perm uint8

// Access permissions.
const (
	// PermR specifies that the memory is readable.
	PermR Perm = 0x4
	// PermW specifies that the memory is writeable.
	PermW Perm = 0x2
	// PermX specifies that the memory is executable.
	PermX Perm = 0x1
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
