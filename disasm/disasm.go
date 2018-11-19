// Package disasm provides general disassembler primitives.
package disasm

import (
	"log"
	"os"
	"sort"

	"github.com/decomp/exp/bin"
	"github.com/mewkiz/pkg/jsonutil"
	"github.com/mewkiz/pkg/osutil"
	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
)

// TODO: Remove loggers once the library matures.

// Loggers.
var (
	// dbg represents a logger with the "disasm:" prefix, which logs debug
	// messages to standard error.
	dbg = log.New(os.Stderr, term.BlueBold("disasm:")+" ", 0)
	// warn represents a logger with the "warning:" prefix, which logs warning
	// messages to standard error.
	warn = log.New(os.Stderr, term.RedBold("warning:")+" ", 0)
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
	Chunks map[bin.Address]map[bin.Address]bool
	// Fragments; sequences of bytes.
	Frags []*Fragment
}

// New creates a new Disasm for accessing the assembly instructions of the given
// binary executable, and the information contained within associated JSON and
// LLVM IR files.
//
// Associated files of the generic disassembler.
//
//    funcs.json
//    blocks.json
//    tables.json
//    chunks.json
//    data.json
func New(file *bin.File) (*Disasm, error) {
	// Prepare generic disassembler.
	dis := &Disasm{
		File:   file,
		Tables: make(map[bin.Address][]bin.Address),
		Chunks: make(map[bin.Address]map[bin.Address]bool),
	}

	// Parse function addresses.
	if err := parseJSON("funcs.json", &dis.FuncAddrs); err != nil {
		return nil, errors.WithStack(err)
	}
	sort.Sort(bin.Addresses(dis.FuncAddrs))

	// Parse basic block addresses.
	if err := parseJSON("blocks.json", &dis.BlockAddrs); err != nil {
		return nil, errors.WithStack(err)
	}
	sort.Sort(bin.Addresses(dis.BlockAddrs))

	// Add entry point function to function and basic block addresses.
	dis.FuncAddrs = bin.InsertAddr(dis.FuncAddrs, dis.File.Entry)
	dis.BlockAddrs = bin.InsertAddr(dis.BlockAddrs, dis.File.Entry)

	// Add export functions to function and basic block addresses.
	for addr := range dis.File.Exports {
		dis.FuncAddrs = bin.InsertAddr(dis.FuncAddrs, addr)
		dis.BlockAddrs = bin.InsertAddr(dis.BlockAddrs, addr)
	}

	// Parse jump table targets.
	if err := parseJSON("tables.json", &dis.Tables); err != nil {
		return nil, errors.WithStack(err)
	}

	// Parse function chunks.
	if err := parseJSON("chunks.json", &dis.Chunks); err != nil {
		return nil, errors.WithStack(err)
	}

	// Compute fragments of the binary; distinct byte sequences of either code or
	// data.
	//
	// Parse data addresses.
	var dataAddrs []bin.Address
	if err := parseJSON("data.json", &dataAddrs); err != nil {
		return nil, errors.WithStack(err)
	}
	// Append basic block addresses to fragments.
	for _, blockAddr := range dis.BlockAddrs {
		frag := &Fragment{
			Addr: blockAddr,
			Kind: KindCode,
		}
		dis.Frags = append(dis.Frags, frag)
	}
	// Append data addresses to fragments.
	for _, dataAddr := range dataAddrs {
		frag := &Fragment{
			Addr: dataAddr,
			Kind: KindData,
		}
		dis.Frags = append(dis.Frags, frag)
	}
	// Sort fragments based on address.
	less := func(i, j int) bool {
		return dis.Frags[i].Addr < dis.Frags[j].Addr
	}
	sort.Slice(dis.Frags, less)

	return dis, nil
}

// IsFunc reports whether the given address is the entry address of a function.
func (dis *Disasm) IsFunc(addr bin.Address) bool {
	less := func(i int) bool {
		return addr <= dis.FuncAddrs[i]
	}
	index := sort.Search(len(dis.FuncAddrs), less)
	if index < len(dis.FuncAddrs) {
		return dis.FuncAddrs[index] == addr
	}
	if _, ok := dis.File.Imports[addr]; ok {
		return true
	}
	return false
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

// ### [ Helper functions ] ####################################################

// parseJSON parses the given JSON file and stores the result into v.
func parseJSON(jsonPath string, v interface{}) error {
	if !osutil.Exists(jsonPath) {
		warn.Printf("unable to locate JSON file %q", jsonPath)
		return nil
	}
	dbg.Printf("parsing: %q", jsonPath)
	return jsonutil.ParseFile(jsonPath, v)
}
