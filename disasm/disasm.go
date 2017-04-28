// Package disasm provides general disassembler primitives.
package disasm

import (
	"github.com/decomp/exp/bin"
)

// A Disasm tracks information required to disassemble a binary executable.
//
// Data should only be written to this structure during initialization. After
// initialization the structure is considered in read-only mode to allow for
// concurrent decoding of functions.
type Disasm struct {
	// Binary executable.
	File *bin.File
	// Function addresses.
	FuncAddrs []bin.Address
	// Basic block addresses.
	BlockAddrs []bin.Address
	// Map from jump table address to target addresses.
	Tables map[bin.Address][]bin.Address
	// Map from basic block address to function address. The basic block is a
	// function chunk and part of a discontinuous function.
	Chunks map[bin.Address]bin.Address
	// Fragments; sequences of bytes.
	Frags []*Fragment
}

// A Fragment represents a sequence of bytes (either code or data).
type Fragment struct {
	// Start address of fragment.
	Addr bin.Address
	// Byte sequence type (code or data).
	Kind FragmentKind
}

// FragmentKind specifies the set of byte sequence types (either code or data).
type FragmentKind uint

// Fragment kinds.
const (
	// The sequence of bytes contains code.
	KindCode = iota + 1
	// The sequence of bytes contains data.
	KindData
)

// New creates a new Disasm for accessing the assembly instructions of the
// given binary executable.
func New(file *bin.File) *Disasm {
	dis := &Disasm{
		File:   file,
		Tables: make(map[bin.Address][]bin.Address),
		Chunks: make(map[bin.Address]bin.Address),
	}
	return dis
}
